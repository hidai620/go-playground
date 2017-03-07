package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(strings.Join([]string{"A", "B","C"}, ","))


	fmt.Println(strings.Index("ABCDE", "A")) // 0
	fmt.Println(strings.Index("ABCDE", "E")) // 4
	fmt.Println(strings.Index("ABCDE", "X")) // -1

	fmt.Println(strings.LastIndex("ABCDECD", "CD")) // 5

	fmt.Println(strings.IndexAny("ABCDECD", "ABC")) // 0

	fmt.Println(strings.LastIndexAny("ABCDECD", "DCB")) // 6 第二引数のいずれかの文字が最後に現れる位置

	fmt.Println("HasPrefix-----------------------------------------") // 2
	fmt.Println(strings.HasPrefix("ABCDEFG", "BC")) // false
	fmt.Println(strings.HasPrefix("ABCDEFG", "AB")) // true

	fmt.Println("HasSuffix-----------------------------------------") // 2
	fmt.Println(strings.HasSuffix("ABCDEFG", "AB")) // false
	fmt.Println(strings.HasSuffix("ABCDEFG", "EFG")) // true


	fmt.Println("Count-----------------------------------------") // 2
	fmt.Println(strings.Count("ABCDEFG ABCDEFG", "ABC")) // 2
	fmt.Println(strings.Count("ABC", "")) // 4


	fmt.Println("Replace-----------------------------------------") // 2
	fmt.Println(strings.Replace("AAAAA", "A", "X", 2))  //XXAAA
	fmt.Println(strings.Replace("AAAAA", "A", "X", -1)) //XXXXX


	fmt.Println("Split-----------------------------------------") // 2

	fmt.Println(len(strings.Split("A,B,C,,", ","))) //5
	fmt.Println(strings.Split("A,B,C,,", ",")) //[A B C  ]
	fmt.Println(strings.SplitAfter("A,B,C,,", ",")) // [A, B, C, , ] セパレーターがついた状態

	fmt.Println(strings.SplitN("A,B,C,,", ",", 2)) //[A B,C,,] 2つに分割する
	fmt.Println(strings.SplitAfterN("A,B,C,,", ",", 2)) //[A, B,C,,] 2つに分割する セパレーターがついた状態

	fmt.Println(len(strings.SplitN("A,B,C,,", ",", -1))) //[A B C   ] -1を指定するとSplitと同じ挙動
	fmt.Println(strings.SplitN("A,B,C,,", ",", -1)) //[A B C   ]


	fmt.Println("To-----------------------------------------") // 2
	fmt.Println(strings.ToLower("ABCDE"))
	fmt.Println(strings.ToUpper("abcde*"))

	fmt.Println("Trim-----------------------------------------") // 2
	fmt.Println(strings.TrimSpace(" abcde "))
	fmt.Println(strings.TrimSpace("　アイウエオ　")) // 全角スペースも除去
	fmt.Println(strings.TrimSpace("\tアイウエオ\n")) // タブ、改行コードも除去

	fmt.Println("Field -----------------------------------------") // 2
	fmt.Println(strings.Fields("a b c d e")) // string配列に分割する
	fmt.Println(strings.Fields("a\nb\nc\nd\ne")) // string配列に分割する
}
