package main

import (
	"fmt"
	"math"
)

var p = fmt.Println


/*
 int32は+,-, 21億まで表現できる。
 int64は+,-, 900京まで表現できる。
 MacBookPro 64bitの環境では、intはint64と同じ最大値、最小値を格納できる。
 32bit環境で、64bitの最大値の書かれたコードをコンパイルしようとするとコンパイルできない。

 uintの場合、符号がないため、マイナスの表現がない。
 その代わり、正の値はintの2倍の数値を表現できる。
 数値型の最大値はmathパケージで参照できる。
 */
func main() {
	var i8 int8 = 127
	var i8_2 int8 = -128
	// var i8_3 int8 = -129 // constant -129 overflows int8
	// var i8_4 int8 = 128 //constant 128 overflows int8

	p(i8)
	p(i8_2)



	var i32 int32 = 2147483647 // +,-, 21億まで表現できる。
	var i32_2 int32 = -2147483648 //
	// var i32_3 int32 = 2147483648 // overflow

	p(i32)
	p(i32_2)

	var i64 int64 = 9223372036854775807
	p(i64)

	var i int = 9223372036854775807
	p(i)

	// 暗黙の型変換をしないのでエラー。 (キャストすると代入できる)
	// i = i64 // cannot use i64 (type int64) as type int in assignment
	i = int(i64) // cannot use i64 (type int64) as type int in assignment


	p(math.MaxInt32)
	p(math.MaxInt64)
}
