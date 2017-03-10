package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

func main() {
	p(5/3)  // 1
	p(5%3)  // 2
	p(-5/3) // -1
	p(-5%3) // -2
	p(5/-3) // -1
	p(5%-3) // 2
	p(-5/-3) // 1
	p(-5%-3) // -2
}
