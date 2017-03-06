package main

import (
	"time"
	"fmt"
)

func main() {
	t, err := time.Parse(time.RFC3339, "2017-03-07T14:30:00Z")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t) // 2017-03-07 14:30:00 +0000 UTC

	t = t.AddDate(0,1,0)
	fmt.Println(t) // 2017-04-07 14:30:00 +0000 UTC


	// "2006-01-02T15:04:05Z07:00"
	format := "2006年01月02日"
	fmt.Println(t.Format(format)) // 2017年04月07日
}