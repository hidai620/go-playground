package main

import (
	"net/url"
	"fmt"
)

func main() {
	url := new(url.URL)

	url.Scheme = "https"
	url.Host = "google.co.jp"
	url.Path = "webhp"

	query := url.Query()
	query.Add("sourceid","chrome-instant")
	query.Add("ion", "1")
	query.Add("espv","2")
	query.Add("ie", "UTF-8")
	query.Add("q","検索")

	url.RawQuery = query.Encode()
	fmt.Println(url)
}
