package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

// append
// スライスの要素の追加関数appendは、可変長引数で値を追加できる。
// 戻り値を受け取らないとコンパイルエラー。
// appendの引数に渡したインスタンスに副作用はない。戻り値だけ値が追加されている。
// 追加要素をスライスで受け取る時、...で可変長引数として展開する。
func main() {
	s := []int{1,2,3}
	p(s)

	s = append(s, 4)
	p(s)

	s = append(s, 5,6,7)
	p(s)

	// 戻り値を取らないappendはコンパイルエラーになる。
	// append(s, 8)

	// 戻り値を別の変数で受けると、
	// appendインスタンス自体は値が変わっていないことがわかる。
	s2 := append(s, 8)
	p(s2) // 1,2,3,4,5,6,7,8
	p(s)  // 1,2,3,4,5,6,7


	// 別のスライスを末尾に追加する。...でスライスの値を可変長引数として展開する。
	s3 := []int{8,9,10}
	s = append(s, s3...)
	p(s)

	// これでも同じ
	//s = append(s, []int{8,9,10}...)


	var b []byte
	b = append(b, "ABCDE"...)
	b = append(b, "FGHIJK"...)
	p(b) //[]byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b}
	p(string(b)) // "ABCDEFGHIJK"
}