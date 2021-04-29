package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aizatto/urlf/url"
	"github.com/pkg/errors"
)

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		println(
			`usage: urlf <format> <url>

Print components of a url 

Commands:
  urlf <url>
    Prints outs the components of a url

  urlf <format> <url...>
    Prints outs the url according to the given format

    Format:
      - s: Scheme
      - u: User
      - U: Password
      - h: Host
      - P: Port
      - p: Path
      - q: Query
      - f: Fragment

    Supports multiple urls

    Example: 
      urlf "h P" http://example.com:443 https://aizatto.com/
`,
		)
		os.Exit(1)
	case 1:
		out, err := urlhelper(args[0])
		if err != nil {
			println(err)
			os.Exit(1)
		}
		fmt.Printf(out)
	default:
		for _, rawurl := range args[1:] {
			out, err := urlf(args[0], rawurl)
			if err != nil {
				println(err)
				os.Exit(1)
			}
			fmt.Printf(out)
			fmt.Print("\n")
		}
	}
}

func urlhelper(urlstr string) (string, error) {
	url, err := url.Parse(urlstr)
	if err != nil {
		return "", errors.Wrapf(err, "unable to parse")
	}

	var builder strings.Builder
	fmt.Fprintf(&builder, "scheme: %s\n", ifempty(url.Scheme))
	fmt.Fprintf(&builder, "username: %s\n", ifempty(url.Username))
	password := url.Password
	fmt.Fprintf(&builder, "password: %s\n", ifempty(password))
	fmt.Fprintf(&builder, "host: %s\n", ifempty(url.Host)) // todo strip port
	fmt.Fprintf(&builder, "port: %s\n", url.Port)          // should display default port
	fmt.Fprintf(&builder, "path: %s\n", ifempty(url.Path))
	fmt.Fprintf(&builder, "query: %s\n", ifempty(url.Query))
	fmt.Fprintf(&builder, "fragment: %s\n", ifempty(url.Fragment))

	return builder.String(), nil
}

func ifempty(value string) string {
	if value == "" {
		return "<empty>"
	}

	return value
}
