package main

import "fmt"

var p = fmt.Println
var pf = fmt.Printf

/*
 16進数で表現したリテラルを標準出力すると10進数で表示される。
 値型の文字に値を入れることで型を変換できる。(キャスト)
 リテラルで定義した型はint型になる。
 */
func main() {
	max := 0xFF
	p(max) // 255
	p(255 == int(max)) //true キャスト

	var i int
	var i64 int64 = 100

	// i = i64 これはエラー
	i = int(i64) // キャストして変換すると値を代入できる。
	p(i) // 100


	num := 1000
	pf("%T", num) // int


}
