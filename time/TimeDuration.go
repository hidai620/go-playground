package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Date(2017, 1, 31, 14, 32, 0, 0, time.Local)
	fmt.Println(t) // 2017-01-31

	t = t.Add(2 * time.Minute) // 2分後
	fmt.Println(t)

	t = t.AddDate(0, 1, 0) // 1ヶ月後
	fmt.Println(t) //  2017-03-03 .. 大分値がおかしい
}
