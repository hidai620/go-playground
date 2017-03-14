package main

import "fmt"

var p = func(x ...interface{}) {
	fmt.Printf("%#v\n", x)
}

type Point struct {
	X, Y int
}

func (point *Point) Render() {
// func (point Point) Render() { // こう定義してもポインタ型から呼び出せる。
	p(point.X, point.Y)

}


// オーバーロードできない
//func (point *Point) Render(format string) { // Duplicate method name
//	p(point.X, point.Y)
//
//}

/*
メソッドはポインタ型に対して定義しても、
値型に定義しても、どちらのインスタンスからでも呼び出せる。
メソッドのオーバーロードはできない
 */
func main() {

	p := Point{X:1, Y:2}
	p.Render()

	p2 := &Point{X:1, Y:2}
	p2.Render()
}
