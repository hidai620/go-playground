package main

import (
	"net/http"
	"net/url"
)

func main() {
	body := url.Values{}
	body.Add("name", "Tom")
	body.Add("id", "1")
	http.PostForm("http://localhost:8080/?foo=bar", body)
}
