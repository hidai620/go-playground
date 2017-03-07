package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseFloat("123.456", 10))
	fmt.Println(strconv.ParseFloat("123.456", 64))
	fmt.Println(strconv.ParseFloat(".456", 64))
	fmt.Println(strconv.ParseFloat("-2", 64))
	fmt.Println(strconv.ParseFloat("12345e8", 64))
	fmt.Println(strconv.ParseFloat("1.00000000001", 32)) // 1
	fmt.Println(strconv.ParseFloat("1.00000000001", 64))

	fmt.Println(strconv.ParseFloat("1E500", 64))  // +Inf strconv.ParseFloat: parsing "1E500": value out of range
	fmt.Println(strconv.ParseFloat("-1E500", 64)) // -Inf strconv.ParseFloat: parsing "-1E500": value out of range
}
