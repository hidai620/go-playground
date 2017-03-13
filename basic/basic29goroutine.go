package main

import (
	"fmt"
	"time"
	"runtime"
)

var p = func(s ...interface{}) {
	fmt.Printf("%s\n", s)
}

// runtimeパッケージは実行環境に関する情報を取得できるメソッドを持つ。
// NumGoroutineで実行しているゴルーチンの数を確認できる。
func main() {
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGroutine: %d\n", runtime.NumGoroutine()) // 1
	fmt.Printf("Version: %s\n", runtime.Version())

	go sub()
	fmt.Printf("NumGroutine: %d\n", runtime.NumGoroutine()) // 2
	//
	//for {
	//	println("main loop")
	//	time.Sleep(1 * time.Second)
	//}

}

func sub() {
	for {
		println("sub loop")
		time.Sleep(5 * time.Second)
	}
}
