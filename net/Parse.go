package main

import (
	"net/url"
	"fmt"
)

func main() {
	url, err := url.Parse("https://mail.google.com/mail/u/0/?tab=mm#inbox")
	if err != nil {
		println(err)
	}

	println(url.Host) // mail.google.com
	println(url.ForceQuery) // false
	println(url.Path) // /mail/u/0
	println(url.RawPath)
	println(url.RawQuery) // tab=mm
	println(url.Scheme) // https
	fmt.Printf("%#v", url.User)
	fmt.Printf("%#v", url.Query()) // https
}
