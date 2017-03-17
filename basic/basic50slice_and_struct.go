package main

import (
	"fmt"
	"time"
)

type Point struct {
	X, Y int
}

var p = func(i interface{}) {
	fmt.Printf("%#v\n", i)
}

/*
構造体を含むスライスを生成する場合、
スライス初期化の際に、内部の構造体も初期化した方が生成時間は一番短い

ただし、最終的に必要なスライスを作成する際には、
初期化したスライスをループしながら、内部の構造体の値を変更していくことになる。
これは他のプログラミング言語を経験した人からみると、可読性がよくない可能性がある。
可読性、速度で見たときにもっとも効率がいいのは、
スライス生成時に必要なスライスのサイズを0、容量を必要な件数で初期化したもの。
容量を指定しない場合と比べて、5倍以上の性能改善がある。

実行結果===================================================================================
        56ミリ秒 スライスをサイズ、容量10000で初期化
        64ミリ秒 スライスをサイズ0、容量10000で初期化し、Pointインスタンスはループで生成しながら追加
       450ミリ秒 スライスをサイズ0、容量0で初期化し、Pointインスタンスはループで生成しながら追加
==========================================================================================

 */
func main() {
	// スライス生成時に同時に、必要な件数分用意し内部の構造体も初期化
	run(makeSlice1, "スライスをサイズ、容量10000で初期化")
	run(makeSlice2, "スライスをサイズ0、容量10000で初期化し、Pointインスタンスはループで生成しながら追加")
	run(makeSlice3, "スライスをサイズ0、容量0で初期化し、Pointインスタンスはループで生成しながら追加")
}

func run(fn func (), description string) {
	start := time.Now().UnixNano()

	fn()

	end := time.Now().UnixNano()

	result := end - start
	fmt.Printf("%10dミリ秒 %s\n", (result/1000), description)
}


const count = 10000

func makeSlice1() {
	points := make([]Point, count, count)

	for i, v := range points {
		v.X = i
		v.Y = i
	}
}

// スライス生成時にサイズ0、容量10000を指定し、
// ループでPointインスタンスを生成しながらスライスに追加
func makeSlice2() {
	points := make([]Point, 0, count)

	for i := 0; i < count; i++ {
		points = append(points, Point{X:i, Y:i})
	}
}

// スライスをサイズ、容量0で作成
// ループでPointインスタンスを生成しながらスライズに追加
func makeSlice3() {
	points := make([]Point,0)

	for i := 0; i < count; i++ {
		points = append(points, Point{X:i, Y:i})
	}
}
