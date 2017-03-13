package main

import "fmt"

var p = func(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}
/*
len=0,   cap=0    初期値
len=1,   cap=1    1追加
len=4,   cap=4    2,3,4追加
len=20,  cap=20   20まで追加
len=100, cap=112  100まで追加
 */
func main() {
	s := make([]int, 0,0)
	p(s)

	s = append(s, 1)
	p(s)

	s = append(s, 2,3,4)
	p(s)

	s = append(s, arrays(s[len(s) -1], 20)...)
	p(s)
	//
	s = append(s, arrays(s[len(s) -1], 100)...)
	p(s)

	fmt.Printf("%#v", s)
}

func arrays(min, max int) []int {
	//fmt.Printf("min:%d, max:%d\n", min, max)
	s := make([]int, 0, max)

	next := min + 1
	for next <= max {
		s = append(s, next)
		next += 1
	}
	//fmt.Printf("%#v, %d \n", s, len(s))
	return s
}