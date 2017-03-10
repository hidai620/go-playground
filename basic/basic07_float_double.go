package main

import "fmt"

var p = fmt.Println

/*
Eで指数を表現する。
1E1 = 1 * (10) = 10
1E2 = 1 * (10 * 10) = 100
 */
func main() {
	p(0.)
	p(10.10)
	p(010.10)
	p(1.e0)
	p(1.E1)   // 1 * 10          = 10
	p(1.e1)   // 1 * 10          = 10
	p(1E1)    // 1 * 10          = 10
	p(1.e+1)  // 1 * 10          = 10
	p(1.e+2)  // 1 * (10 * 10)   = 100
	p(1E2)    // 1 * (10 * 10)   = 100
	p(1.e-1)  // 1 * 0.1         = 0.1
	p(1.e-2)  // 1 * (0.1 * 0.1) = 0.01
}
