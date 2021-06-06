package url

import (
	url2 "net/url"
	"strings"

	"github.com/pkg/errors"
)

type URL struct {
	raw      string
	Scheme   string
	Username string
	Password string
	Host     string
	Port     string
	Path     string
	Query    string
	Fragment string
}

func Parse(rawurl string) (*URL, error) {
	oldurl, err := url2.Parse(rawurl)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse")
	}

	newurl := &URL{}
	newurl.raw = rawurl

	newurl.Scheme = oldurl.Scheme
	newurl.Username = oldurl.User.Username()
	newurl.Password, _ = oldurl.User.Password()

	host := oldurl.Host
	index := strings.Index(host, ":")
	if index != -1 {
		host = host[:index]
	}
	newurl.Host = host

	newurl.Port = oldurl.Port()
	if newurl.Port == "" {
		switch oldurl.Scheme {
		case "ftp":
			newurl.Port = "21"
		case "ssh":
			newurl.Port = "22"
		case "sftp":
			newurl.Port = "22"
		case "http":
			newurl.Port = "80"
		case "https":
			newurl.Port = "443"
		}
	}

	newurl.Path = oldurl.Path
	if newurl.Path == "" {
		newurl.Path = "/"
	}
	newurl.Query = oldurl.RawQuery
	newurl.Fragment = oldurl.Fragment

	return newurl, nil
}
