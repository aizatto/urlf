package main

import (
	"strings"

	"github.com/aizatto/urlf/url"
	"github.com/pkg/errors"
)

func urlf(format, rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", errors.Wrapf(err, "unable to parse")
	}

	formatlen := len(format)

	var builder strings.Builder
	for i := 0; i < formatlen; i++ {
		// organized in the order of a URI scheme
		switch format[i] {
		case 's':
			builder.WriteString(url.Scheme)
		case 'u':
			builder.WriteString(url.Username)
		case 'U':
			builder.WriteString(url.Password)
		case 'h':
			builder.WriteString(url.Host)
		case 'P':
			builder.WriteString(url.Port)
		case 'p':
			builder.WriteString(url.Path)
		case 'q':
			builder.WriteString(url.Query)
		case 'f':
			builder.WriteString(url.Fragment)
		default:
			builder.WriteByte(' ')
		}
	}

	return builder.String(), err
}
