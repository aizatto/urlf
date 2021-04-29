package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func musturlf(format, url string) string {
	out, err := urlf(format, url)
	if err != nil {
		panic(err)
	}

	return out
}

func TestUrlf(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		"example.com 80",
		musturlf("h P", "http://example.com"),
		"Must match",
	)

	assert.Equal(
		"example.com 443 /",
		musturlf("h P p", "https://example.com"),
		"Must match",
	)

	assert.Equal(
		"http example.com 443",
		musturlf("s h P", "http://example.com:443"),
		"Must match",
	)

	assert.Equal(
		"http user password example.com 443 /path query fragment",
		musturlf("s u U h P p q f", "http://user:password@example.com:420/path?query#fragment"),
		"Must match",
	)
}
