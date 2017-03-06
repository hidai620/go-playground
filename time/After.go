package main

import (
	"time"
	"fmt"
)

/**
 * 1秒ごとにタスクを実行し、5秒後にタイムアウトさせる。
 */
func main() {
	fmt.Println(time.Now())
	timeoutCh := time.After(5 * time.Second)

	ch := time.Tick(1 * time.Second)

	// L:
	for {
		select {
		case t := <-ch:
			fmt.Println("tick:", t)
		case m := <-timeoutCh:
			fmt.Println("end", m)
			return
			// break L これでも抜けられる
		}
	}
}
