package main

import "fmt"

var p = fmt.Println
var pf = fmt.Printf

/**
 byte型はuint8の別名で定義されている。
 byte型はuint8で表現できるだけの数字でしかない。
 */
func main() {
	var b byte = byte(10)
	p(b)
	pf("%T\n", b) // uint8

	p("最小", byte(0))
	p("最大", byte(255))
	// p(byte(256)) // コンパイルエラー
}
