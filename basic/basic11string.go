package main

import "fmt"

var p = func(s string) {
	fmt.Printf("%v\n", s)
}

func main() {
	p("\xff\uff00")
	p("\u65e5")
	p("ab\nc")
	p(`ab\nc`)
}
