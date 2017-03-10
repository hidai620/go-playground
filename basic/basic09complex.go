package main

import "fmt"

var p = fmt.Println

func main() {
	p(1.0 + 3i)
	p(complex(1.0,3))


	p(0i)    // 0+0i
	p(0.i)   // 0+0i
	p(2.781i)// 0+2.781i

	c := (1.0 + 3i)
	p(c) // 1+3i
	p(real(c)) // 1 // floatで返る
	p(imag(c)) // 3 // floatで返る
}
