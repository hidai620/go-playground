package main

import "fmt"

var p = func(s []int) {
	fmt.Printf("len=%d, cap=%d, %#v \n", len(s), cap(s), s)
}

/*
スライスのコピー先の変数に、コピー対象の変数を上書きする。
コピーできる要素の数は、コピー先のスライスの長さに依存する。
コピー先以上のスライスをコピーしようとした場合、結果がコピー先以上の長さに拡張されることはない。

完全スライス式は容量まで指定する。
 */
func main() {
	s1 := []int {1,2,3,4,5}
	s2 := []int {10,11}

	n := copy(s1, s2) // s1に副作用のある関数
	p(s1) // len=5, cap=5, []int{10, 11, 3, 4, 5}
	p(s2) // len=2, cap=2, []int{10, 11}
	fmt.Printf("%#v", n) // 2


	//
	s3 := make([]byte, 10, 100)
	copy(s3, "abcde")
	fmt.Printf("%#v\n", s3)
	fmt.Printf("%#v\n", string(s3))



	s4 := make([]int, 3,10)
	s4[0] = 1
	s4[1] = 2
	s4[2] = 3

	p(s4)

	// コピーする対象は要素数1から配列の3つ目までの2つ。
	// コピーする容量の設定は要素数1から配列の7つ目までの6つ。
	s5 := s4[1:3:7]
	p(s5) // len=2, cap=6, []int{2, 3}
}
