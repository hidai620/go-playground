package main

import (
	"time"
	"fmt"
)

func main() {
	// 2017年3月6日 14時32分0秒0ナノ秒
	t := time.Date(2017, 3, 6, 14, 32, 0, 0, time.Local)
	fmt.Println(t)

	fmt.Println("Year", t.Year())
	fmt.Println("YearDay", t.YearDay()) // 1月1日から何日目か
	fmt.Println("Month", t.Month()) // 英字
	fmt.Println("Month", int(t.Month())) // 数字
	fmt.Println("Weekday", t.Weekday())
	year, month, day := t.Date()
	fmt.Println("Date", year, month, day)
	fmt.Println("Day", t.Day())
	fmt.Println("Hour", t.Hour())
	fmt.Println("Minute", t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Local())
}
