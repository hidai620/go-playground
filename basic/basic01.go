package main

import "fmt"

type user struct {
	id int
	name string
}

var p = func (v ...interface{}) {
	fmt.Printf("%#v\n", v)
}

/**
 型は初期化しなくても、
 参照した時にデフォルトの値で初期化される。
 */
func main() {
	var i int
	var s string


	p(i)
	p(s)
	p(user{})
}

