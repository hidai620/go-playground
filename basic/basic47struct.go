package main

import "fmt"

var p = func(x interface{}) {
	fmt.Printf("%#v\n", x)
}

type X0 struct {
	X0
}// 自身の型を持ってもいけない。再帰的と見なされる。

type X1 struct {
	X2
}

// 再帰的な構造体の定義はエラーになる
type X2 struct {
	X1
} // invalid recursive type X2

func main() {

	var x X1
	p(x)
}
