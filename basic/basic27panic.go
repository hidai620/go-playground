package main

import "fmt"

var p = func(s ...interface{}) {
	fmt.Printf("%s\n", s)
}

// パニックを発生した場合、deferの登録処理は実行される。
func main() {
	p("main start")
	defer fmt.Println("defer 1")

	panic("panic!!!")

	p("main end")
}
