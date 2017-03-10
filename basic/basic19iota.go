package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

const (
	a = 1 + iota
	b
	c
	d = 100
	e = 1 + iota
)

/*
定数はiotaで表現できる。
iotaの値は定数ないの数で決まる。
 */
func main() {
	p(a)
	p(b)
	p(c)
	p(d)
	p(e) // 5
}
