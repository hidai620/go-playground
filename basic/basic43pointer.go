package main

import "fmt"

var p = fmt.Println

/*
# ポインタ型
ポインタ型で定義された変数の初期値はnilになる。
アドレス演算子&を値型の変数に使うことで、その変数のポインタ型を取得できる。
ポインタ型に値を直接代入することはできない。
ポインタ型からそのアドレスの先の値を参照するためには「*」演算子をポイン型変数の前に置き参照する。（デリファレンス)
nilのポインタ型変数からデリファレンスするとパニックが起きる。
値型を関数の引数に渡すと、値のコピーが発生する。
ポインタ型を関数の引数に渡すことで、値はコピーされず、同じメモリ上のアドレスを参照できる。

# 配列のポインタ型の操作
// C言語由来の書き方 デリファレンスでnumsのさすアドレスの本体を参照
// n := (*nums)[i]
// (*nums)[i] = n*n

// Goではこう書くだけでポインタ型のデリファレンスを参照しているとコンパイラに見なされる。
// ポインタのさすメモリ上の値本体を見る書き方をしなくてもいい。
// （デリファレンスが不要になりソースが見やすくなる)
n := nums[i]
nums[i] = n*n


# 文字(string)のポインタ
stringは不変(immutable)
stringのデリファレンスを取得することができるが、その値を変更することはできない。

stringの結合はその度に値がコピーされ、新しいメモリの割り当てが起きる。
stringを関数の引数に渡しても値のコピーは置きない。
string型は構造体で内部にポインタを持っている。
関数の引数にstringを渡した場合、コピーされるのは、stringの構造体内部のポインターと文字列長さのみ。

 */
func main() {
	var i *int

	p(i) // nil
	fmt.Printf("%#v, address=%p \n", i , i)

	// i = 10 // cannot use 10 (type int) as type *int in assignment

	// nilの状態のポインタ型からデリファレンスすることはできない。
	// *i = 5 // panic: runtime error: invalid memory address or nil pointer dereference


	x := 10
	pointer := &x
	p(pointer, *pointer) // 10
	x = 20
	p(pointer, *pointer) // 20

	i = pointer

	*i = 10
	p(i, *i) // 10
	p(pointer, *pointer) // 10
	fmt.Printf("%#v, address=%p \n", i , i)


	nums := &[3]int{1,2,3}
	pow(nums)
	p(nums) // &[1 4 9]


	// 組み込み関数、スライス式、forを利用する際もデリファレンス不要
	p(len(nums))
	p(cap(nums))
	p(nums[1:2])

	for i, v := range nums {
		p(i,v)
	}


	// 文字列のさすアドレスの参照、参照先の値
	s := "a"
	ps := &s
	*ps = "b"
	p(s)
	p(s[0])
	// 文字の部分参照をポインタ型で得ることはできない。
	// p(&s[0]) // cannot take the address of s[0]


}

func pow(nums *[3]int) {
	for i := 0; i < len(nums); i++ {
		// C言語由来の書き方 デリファレンスでnumsのさすアドレスの本体を参照
		//n := (*nums)[i]
		//(*nums)[i] = n*n

		// Goではこう書くだけでポインタ型のデリファレンスを参照しているとコンパイラに見なされる。
		// ポインタのさすメモリ上の値本体を見る書き方をしなくてもいい。
		// （デリファレンスが不要になりソースが見やすくなる)
		n := nums[i]
		nums[i] = n*n
	}
}
