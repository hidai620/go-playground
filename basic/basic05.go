package main

import (
	"fmt"
	"math"
)

var p = fmt.Println
var pf = fmt.Printf

/**
 型変換や計算した値の桁あふれは、overflowでエラーにならない。
 ラップアラウンドされる。
 */
func main() {
	b := byte(255)
	/*
	2進数で表現した255が
	  1111 1111
	1加算されたことで一桁繰り上がって
	  0000 0000
	になり、0が表示される。
	 */
	p(b + 1) // 0


	// 型え表現できる数値がループするようなラップアラウンドが起こると、
	// 40億 + 4億が 1億弱の結果になってしまう。105032704
	var i1 uint32 = 4000000000
	var i2 uint32 = 400000000
	p(i1 + i2) // 105032704

	p(isSafeCalculation(i1, i2)) //false
}

func isSafeCalculation(a, b uint32) bool {
	// 型が表現できる最大値から足す対象の数値を引いた値が、
	// 足そうとしている数値に安全に足せる数値。
	var safeValue = math.MaxUint32 - a

	// 安全な値が足そうとしている数値より小さければ、ラップアラウンドは起きない。
	return safeValue < b
}
