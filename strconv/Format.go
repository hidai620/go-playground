package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.FormatBool(true))

	fmt.Println(strconv.FormatInt(123456, 10)) // 123456
	fmt.Println(strconv.FormatInt(123456, 16)) // 1e240
	fmt.Println(strconv.FormatInt(123456, 2))  // 11110001001000000

	fmt.Println(strconv.FormatFloat(123.456, 'E', -1, 64)) // 1.23456E+02
	fmt.Println(strconv.FormatFloat(123.456, 'e', 2, 64))  // 1.23e+02
	fmt.Println(strconv.FormatFloat(123.456, 'f', -1, 64)) // 123.456 実数で表現
	fmt.Println(strconv.FormatFloat(123.456, 'f', 2, 64))  // 123.45   実数で小数点２位までで表現

}
