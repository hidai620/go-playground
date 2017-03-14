package main

import "fmt"

var p = func(x interface{}) {
	fmt.Printf("%#v\n", x)
}


type Point struct {
	X, Y int
}

// プロパティ名を省略すると、プロパティ名が型名を利用される。
type X struct {
	int
	float64
}

// アンダースコアを指定すると無名フィールドになる。
// 初期化もアクセスもできない。
// メモリ調整などの高度な処理に利用する。
type Y struct {
	_ int
}

/*
## 構造体は値型
変数を宣言すると、初期化しなくても、その変数の型の初期値で初期化される。

## 構造体の値を生成するときの注意
中カッコで生成するときは必ずプロパティ名を指定する。
もし、プロパティの順番に変更があった時、生成する際に値がずれて入ってしまうため。
プロパティ名を指定することで、順序の変更の影響を受けなくなる。

 */
func main() {

	var point Point
	p(point) // main.Point{X:0, Y:0}

	//point2 := Point{1,2} // プロパティの順序変更で値がずれてしまう
	point2 := Point{X:1,Y:2}
	p(point2)

	// プロパティ名省略
	var x X
	p(x) // main.X{int:0, float64:0}

	var y Y
	p(y) // main.Y{_:0}

}
