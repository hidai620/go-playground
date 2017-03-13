package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v \n", s)
}

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"
	p(m) // map[int]string{86:"China", 1:"US", 81:"Japan"}

	// 型無し浮動小数点をキーに使うと、float64で表現できる範囲に丸まってしまう
	n := make(map[float64]string)
	n[1.0000000000000001] = "a"
	n[1.0000000000000002] = "b"
	p(n) // map[float64]string{1:"a", 1.0000000000000002:"b"}

	users := map[int] string {
		1: "Tome",
		2: "Jane",
	}
	p(users)


	// 配列を含むマップはシンタックスシュガーで簡単にかける
	arrays := map[int][] int {
		//1: []int{1,2,3},
		//2: []int{4,5,6},
		1: {1,2,3}, // 上に書いたものと同じものを簡略的にかける
		2: {4,5,6},
	}
	p(arrays)

	// mapを要素に含む場合も同様に簡略表記がある。
	maps := map[int]map[int]string {
		//1: map[int]string {1 : "A", 2: "B"},
		//2: map[int]string {1 : "a", 2: "b"},
		1: {1 : "A", 2: "B"},
		2: {1 : "a", 2: "b"},
	}
	p(maps)
}
