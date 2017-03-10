package main

import (
	"fmt"
	"math"
)

var p = fmt.Println

/*
リテラルでゼロ除算するとエラーで落ちるが、
変数内に格納されている値でゼロ除算してもエラーにはならず、正の無限大、負の無限大の数になる。
 */
func main() {
	p(math.MaxFloat32)
	p(math.SmallestNonzeroFloat32)
	p(math.MaxFloat64)
	p(math.SmallestNonzeroFloat64)

	f64 := 1.0
	fmt.Printf("%#T", f64) // float64

	z := 0.0
	a := 1.0 / z // エラーにならない。正の無限大
	b := -1.0 / z // エラーにならない。負の無限大
	p(a)
	p(b)
	p(z/z) // NaN

	//c := 1.0/0.0 //  error division by zero
	//p(c)
}
