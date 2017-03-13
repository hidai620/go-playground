package main

import "fmt"

var p = func(s ...interface{}) {
	fmt.Printf("%s\n", s)
}

// deferに登録された関数は登録の逆順に実行される。
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")


	fmt.Println("main start")
	fmt.Println("main end")
}
