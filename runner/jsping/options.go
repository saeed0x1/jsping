package jsping

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	InputFile string
	u         string
	Workers   int
	Timeout   int
	stdin     bool
	UserAgent string
	JSONLines bool
	output    string
	Cookie    string
}

func ParseOptions() *Options {
	opts := &Options{}
	flag.StringVar(&opts.InputFile, "f", "", "Provide a file containing URLS")
	flag.StringVar(&opts.UserAgent, "ua", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.67", "User-Agent to send in requests")
	flag.StringVar(&opts.u, "url", "", "Set the URL of the site")
	flag.StringVar(&opts.Cookie, "cookie", "", "Cookie to send in requests")
	flag.StringVar(&opts.output, "o", "", "Write the output in a file")
	flag.IntVar(&opts.Workers, "c", 10, "Number of concurrent requests to send")
	flag.IntVar(&opts.Timeout, "t", 15, "Timeout (in seconds) for http client")
	showVersion := flag.Bool("version", false, "Show version number")
	flag.BoolVar(&opts.stdin, "stdin", false, "Set to take standard input")
	flag.BoolVar(&opts.JSONLines, "json", false, "Output in Json format")
	flag.Parse()
	if *showVersion {
		fmt.Printf("jsping version: %s\n", version)
		os.Exit(0)
	}
	return opts
}

func PrintUsage() {
	flag.Usage()
}
