package main

import (
	"time"
	"fmt"
)

func main() {
	var err error

	t, err := time.Parse(time.RFC3339, "2017-03-07 14:00:30") // error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	format := "2006-01-02T15:04:05"
	t, err = time.Parse(format, "2017-03-07T14:00:30")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t) // 2017-03-07 14:00:30 +0000 UTC

	format = "2006-01-02"
	t, err = time.Parse(format, "2017-03-07")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t) // 2017-03-07 00:00:00 +0000 UTC
}
