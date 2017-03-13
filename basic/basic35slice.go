package main

import (
	"fmt"
)


var p = func(s []int) {
	fmt.Printf("len=%d, cap=%d, %#v \n", len(s), cap(s), s)
}

/*
古典的なfor文とrangeのfor文でスライス自体の参照の仕方が違う。

古典的なfor文は、ブロックの中でループ中のスライスを参照すると参照としてアクセスする。
ループ中のスライスに追加すると、無限ループになる可能性がある。

rangeのfor文は、ブロックの中でループ中のスライスを参照しても値渡しのようにアクセスする。
ループ中のスライスの変数の参照を変更しても、for文の外でスライスの変数の参照に影響しない。

しかしrangeのforでもスライスの要素にアクセスする時は参照渡しのようにアクセスする。
rangeのforの内部で、スライスの要素を変更すると、for文の外のスライスの参照に影響する。
 */
func main() {
	nums := []int{1,2,3}

	// for文の中で値を追加しても追加前の状態しかループは回らない。
	for _, n := range nums {
		fmt.Println(n)
		max := nums[len(nums) -1] + 1
		nums = append(nums, max)
	}
	p(nums)

	// for文の中でスライスに値を追加すると無限ループに陥る例
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		max := nums[len(nums) -1] + 1
		nums = append(nums, max)

		// ブレイクしないと無限ループになる
		if nums[i] > 4 {
			break
		}
	}
	p(nums)

	// スライス自体の参照は値として扱われるが、
	// スライスの中は参照渡しされているため、要素の値を変更すると、ループの外に影響する。
	for i, n := range nums {
		nums[i] = n*n
	}
	p(nums)
}
