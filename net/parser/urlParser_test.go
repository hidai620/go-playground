package parser

import (
	"testing"
	"net/url"
)

func TestParseUrl(t *testing.T) {
	url, err := url.Parse("http://example.com/auth")
	if err != nil {
		t.Error(err)
	}

	url, err = parseURL(url, "login")
	if err != nil {
		t.Error(err)
	}

	path := "/auth/login"
	if url.Path != path {
		t.Errorf("Expected url path to be <%v>, got <%v>", path, url.Path)
	}
}
