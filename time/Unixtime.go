package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Date(2017, 3, 6, 14, 32, 0, 0, time.Local)
	fmt.Println(t) // 2017-03-06 14:32:00 +0900 JST

	unixTime := t.Unix()
	fmt.Println(unixTime) // 1488778320

	t2 := time.Unix(unixTime, 0)
	fmt.Println(t2) // 2017-03-06 14:32:00 +0900 JST
}
