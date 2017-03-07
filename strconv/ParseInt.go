package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(strconv.ParseInt("123", 10, 0))
	fmt.Println(strconv.ParseInt("-1", 10, 0))
	fmt.Println(strconv.ParseInt("740", 10, 0))
	fmt.Println(strconv.ParseInt("FF", 16, 0)) // 16進数の文字を10進数の数値に変換

	fmt.Println("-----------------------------------------")
	fmt.Println(strconv.ParseUint("123", 10, 0))
	fmt.Println(strconv.ParseUint("-1",  10, 0)) // error strconv.ParseUint: parsing "-1": invalid syntax
	fmt.Println(strconv.ParseUint("740", 10, 0))
	fmt.Println(strconv.ParseUint("FF",  16, 0)) // 16進数の文字を10進数の数値に変換
	fmt.Println("-----------------------------------------")

	fmt.Println(strconv.ParseUint("123", 0, 0))
	fmt.Println(strconv.ParseUint("-1",  0, 0))   // 0 error strconv.ParseUint: parsing "-1": invalid syntax
	fmt.Println(strconv.ParseUint("740", 0, 0))
	fmt.Println(strconv.ParseUint("FF",  0, 0))   // 0 strconv.ParseUint: parsing "FF": invalid syntax
	fmt.Println(strconv.ParseUint("0xFF",  0, 0)) // 16進数の文字を10進数の数値に変換
	fmt.Println("-----------------------------------------")
	fmt.Println(strconv.ParseInt("9223372036854775807", 10, 0)) // intの最大値
	fmt.Println(strconv.ParseInt("99999999999999999999999", 10, 0)) // 9223372036854775807 strconv.ParseInt: parsing "99999999999999999999999": value out of range
}
