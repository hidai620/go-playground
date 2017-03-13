package main

import "fmt"

var p = func(s ...interface{}) {
	fmt.Printf("%s\n", s)
}

// goto文はラベルをつけて同じ関数内のみ移動できる。
func main() {
	p("1")

	goto label3
	p("2")

	label3:
	p("3")

	// 別の関数に飛ぶことはできない
	// goto label4
}

func print(s string) {
	// label4:
	fmt.Println(s)
}
