package main

import "fmt"

var p = func(s ...interface{}) {
	fmt.Printf("%s\n", s)
}

// case 節にはbool型の式を書くことができる。
// この時switch節には値を与えない。
// 省略すると、内部的にtrueが与えられたと見なされる。
// switch節にはどんな値を与えても実行される。
func main() {
	n := 3
	switch  {
	case n < 4:
		p("n is smaller than 4")
	case n >= 4:
		p("n is 4 or bigger than 4")
	}

	m := 4
	switch false { // falseと書いても実行される。
	case m < 4:
		p("m is smaller than 4")
	case m >= 4:
		p("m is 4 or bigger than 4")
	}
}
