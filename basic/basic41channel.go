package main

import (
	"time"
)

/*
チャネルをfor rangeのループで受信すると、チャネルの値を受信し続ける。
チャネルがクローズされれば、ループを終了する。
クローズされなければ、ループを抜けることができない。

 */
func main() {
	ch := make(chan int)

	go func() {
		count := 1
		for {
			if count > 5 {
				continue
			}
			time.Sleep(1 * time.Second)
			value := int(time.Now().Unix() )
			ch <- value
			count += 1
		}
		//close(ch)
		// ch <- 1 クローズされたチャンネルに値を送信するとランタイムパニックになる。

		//fmt.Println("ch closed")
	}()


	for v := range ch {
		println(v)
	}

}
