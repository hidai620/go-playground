package parser

import (
	"net/url"
	"path"
)

// parseURL parses an URL from the net.URL and string path.
func parseURL(url *url.URL, urlPath string) (*url.URL, error) {
       return url.Parse(path.Join(url.Path, urlPath))
}
