package main

import "fmt"

var p = fmt.Println

func main() {
	p('あ')
	fmt.Printf("%T", 'あ')
	p('い')
	p('う')

	p('朝')
	p('\t')   // 9 タブ
	p('\a')   // 7 ベル
	p('\'')   // 39 シングルクオート
	p('\000') //  \  = 8進数 3桁
	p('\007')
	p('\007')
	p('\xff') // 255 // \x = 16進数2桁
	p('\u12e4') //      \u = 16進数4桁
	p('\U000012e4') //  \U = 16進数8桁
}
