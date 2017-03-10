package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

/*
interface型は全ての値を代入できる。
 */
func main() {
	var v interface{}
	p(v) // nil

	v = "a"
	v = 1
	v = .0
	v = [5]int{}

	// interface型に入っているintはそのまま計算できない。
	// アサーションで元の型に戻すと、計算できる。
	var x, y interface{}
	x = 1
	y = 2
	// p(x + y) // invalid operation: x + y (operator + not defined on interface)
	p(x.(int) + y.(int))
}
