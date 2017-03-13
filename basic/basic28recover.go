package main

import "fmt"

var p = func(s ...interface{}) {
	fmt.Printf("%s\n", s)
}

// recoverでpanicに渡された値を取得できる。
// recoverの戻り値で、defer 内部でpanicが実行されたか判断できる。
// panicはJavaで言うところのOutOfMemoryErrorと同じレイヤーに相当する。
// 基本的にpanicを起こすようなアプリケーション設計にしてはいけない。
func main() {
	p("main start")
	defer func () {
		reason := recover()
		if reason != nil {
			fmt.Println("reason:", reason)
		}
	}()

	panic("serviceX is out of service")

	p("main end")
}
