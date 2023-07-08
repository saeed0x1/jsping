package jsping

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const version = `1.0.0`

type JSPing struct {
	client *http.Client
	opts   *Options
}

func New(opts *Options) *JSPing {
	c := &http.Client{
		Timeout:   time.Duration(opts.Timeout) * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	return &JSPing{client: c, opts: opts}
}
func (s *JSPing) Run() error {
	// setup input
	var input *os.File
	var url_input string
	var err error
	// input url
	if s.opts.u != "" {
		url_input = s.opts.u
	} else if s.opts.InputFile != "" {
		// file name
		input, err = os.Open(s.opts.InputFile)
		if err != nil {
			return fmt.Errorf("could not open input file: %s", err)
		}
		defer input.Close()
	} else if s.opts.stdin {
		input = os.Stdin
	} else {
		PrintUsage()
		print("\nuse -stdin flag to take standard input")

	}

	// init channels
	urls := make(chan string)
	results := make(chan string)

	// start workers
	var w sync.WaitGroup
	for i := 0; i < s.opts.Workers; i++ {
		w.Add(1)
		go func() {
			s.fetch(urls, results)
			w.Done()
		}()
	}
	// setup output
	var out sync.WaitGroup
	out.Add(1)
	go func() {
		if s.opts.output == "" {
			for result := range results {
				fmt.Println(result)
			}
		} else {
			file, err := os.OpenFile(s.opts.output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			for result := range results {
				fmt.Println(result)
				data := result + "\n"
				_, err := file.WriteString(data)
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}
		}

		out.Done()
	}()

	if url_input != "" {
		// fmt.Println(url_input)
		parsedURL, err := url.Parse(url_input)
		// print(parsedURL.String())
		if err == nil {
			urls <- parsedURL.String()
		}
	}

	scan := bufio.NewScanner(input)
	for scan.Scan() {
		u := scan.Text()
		if u != "" {
			parsedURL, err := url.Parse(u)
			if err == nil {
				urls <- parsedURL.String()
			}
		}
	}
	close(urls)
	w.Wait()
	close(results)
	out.Wait()
	return nil
}

func (s *JSPing) fetch(urls <-chan string, results chan string) {
	domainMap := make(map[string][]string)

	for u := range urls {
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			continue
		}
		if s.opts.UserAgent != "" {
			req.Header.Add("User-Agent", s.opts.UserAgent)
		}
		if s.opts.Cookie != "" {
			req.Header.Set("Cookie", s.opts.Cookie)
		}
		resp, err := s.client.Do(req)
		if err != nil {
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			continue
		}
		uParsed, err := url.Parse(u)
		if err != nil {
			log.Fatalf("error parsing url: %v", err)
			return
		}

		doc.Find("script[src], div[data-script-src]").Each(func(index int, s *goquery.Selection) {
			jsURL := s.AttrOr("src", s.AttrOr("data-script-src", ""))
			if jsURL == "" {
				return
			}

			if !strings.HasPrefix(jsURL, "http://") && !strings.HasPrefix(jsURL, "https://") {
				if strings.HasPrefix(jsURL, "//") {
					jsURL = fmt.Sprintf("%s:%s", uParsed.Scheme, jsURL)
				} else if strings.HasPrefix(jsURL, "/") {
					jsURL = fmt.Sprintf("%s://%s%s", uParsed.Scheme, uParsed.Host, jsURL)
				} else {
					jsURL = fmt.Sprintf("%s://%s/%s", uParsed.Scheme, uParsed.Host, jsURL)
				}
			}

			domain := getDomain(uParsed)
			domainMap[domain] = append(domainMap[domain], jsURL)
		})
	}

	var output []map[string][]string
	for domain, jsURLs := range domainMap {
		if s.opts.JSONLines {
			output = append(output, map[string][]string{domain: jsURLs})
			jsonOutput, _ := json.Marshal(output)
			results <- string(jsonOutput)
		} else {
			for _, jsURL := range jsURLs {
				results <- jsURL
			}
		}
	}
}

func getDomain(u *url.URL) string {
	return u.Scheme + "://" + u.Host
}
