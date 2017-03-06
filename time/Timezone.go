package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Date(2017, 3, 6, 14, 32, 0, 0, time.Local)
	fmt.Println(t) // 2017-03-06 14:32:00 +0900 JST

	fmt.Println(t.UTC()) // 2017-03-06 05:32:00 +0000 UTC
}
