package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(strconv.ParseBool("true")) // true
	fmt.Println(strconv.ParseBool("1"))    // true
	fmt.Println(strconv.ParseBool("t"))    // true
	fmt.Println(strconv.ParseBool("T"))    // true
	fmt.Println(strconv.ParseBool("X"))    // error
	fmt.Println(strconv.ParseBool("True")) // true
	fmt.Println(strconv.ParseBool("true")) // true
	fmt.Println(strconv.ParseBool("tRue")) // error
	fmt.Println(strconv.ParseBool("0")) // false
	fmt.Println(strconv.ParseBool("false")) // false
	fmt.Println(strconv.ParseBool("FALSE")) // false
	fmt.Println(strconv.ParseBool("False")) // false
	fmt.Println(strconv.ParseBool("fAlse")) // error
	fmt.Println(strconv.ParseBool("F")) // false
}
