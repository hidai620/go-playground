package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

// 引数の無視
func do(_, _ int) {
	return
}

func main() {
	do(1,2)
}
