package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		num int
		message string
		x bool
	)

	flag.IntVar(&num, "n", 0, "数値")
	flag.StringVar(&message, "m", "", "最大値")
	flag.BoolVar(&x, "x", false, "フラグ")
	flag.Parse()


	fmt.Println(num)
	fmt.Println(message)
	fmt.Println(x)
}
