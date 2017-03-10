package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

// 定義を省略すると、事前に定義した値を同じ値で初期化される。
// 全ての定義を省略するとコンパイルエラーになる。
// 程数値は計算結果で初期化できる。
const (
	X = 1
	Y
	Z
	A = X + Z
	B bool = Y == Z

	// 定義自体はできるが、使用した際にint型に代入した際にエラーになる。
	C = 999999999999999999999999999999999999999999999999999999
)

func main() {
	p(X)
	p(Y)
	p(Z)
	p(A)
	p(B)
	// p(C + 1)
}