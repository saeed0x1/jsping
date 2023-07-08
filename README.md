# jsping

jsping gathers javascript files from provided url, list of URLS or subdomains. Analysing javascript files give you a lot of insights about your targets.

To see the true potential of this tool pair this tool with [gau](https://github.com/lc/gau) and then [https://github.com/GerbenJavado/LinkFinder](https://github.com/GerbenJavado/LinkFinder)

# Resources
- [Usage](#usage)
- [Installation](#installation)

## Usage:
Examples:
```bash
$ cat urls.txt | jsping -stdin
$ jsping -f urls.txt
$ cat hosts.txt | gau | jsping -stdin
$ jsping -url https://example.com
```

To display the help for the tool use the `-h` flag:

```bash
$ jsping -h
```

| Flag | Description | Example |
|------|-------------|---------|
| `-c` | Number of concurrent requests to send | `jsping -c 40` |
| `-f` | Input file containing URLS | `jsping -f urls.txt` |
| `-t` | Timeout (in seconds) for http client (default 15) | `jsping -t 20` |
| `-ua` | User-Agent to send in requests | `jsping -ua "Chrome..."` |
| `-url` | Take single url as input | `jsping -url https://example.com` |
|`-stdin`| Take standard input from terminal | `cat urls.txt | jsping -stdin`
|`-json`| Output in json format | `cat urls.txt | jsping -stdin -json`
|`-cookie`| Set the cookie | `cat urls.txt | jsping -stdin -cookie "Cookie:..."`
|`-o`| Write the output in a file | `cat urls.txt | jsping -stdin -o output.txt`
| `-version` | Show version number | `jsping -version"` |


## Installation
### From Source:

```
$ GO111MODULE=on go install github.com/saeed0x1/jsping@latest
```

### From Binary
You can download the pre-built [binaries](https://github.com/saeed0x1/jsping/releases/) from the releases page and then move them into your $PATH.

```
$ mv jsping /usr/bin/jsping
```

## Useful?

<a href="https://bmc.link/saeed0x1" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
