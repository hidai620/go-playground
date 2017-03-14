package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v \n", s)
}


/*
送受信チャネルは、送信専用、受信専用チャネルとして宣言された変数に代入できる。
代入してもPrintfでタイプとして参照した場合の型は変わらない。
 */
func main() {
	var ch chan int // 送受信
	var sendCh <-chan int // 送信専用
	var receiveCh chan<- int // 受信専用

	p(ch)        // (chan int)(nil)
	p(sendCh)    // (<-chan int)(nil)
	p(receiveCh) // (chan<- int)(nil)

	sendCh = ch
	receiveCh = ch
	p(sendCh)    // (<-chan int)(nil)
	p(receiveCh) // (chan<- int)(nil)
}
