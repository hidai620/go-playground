package main

import (
	"time"
	"fmt"
)

/*
 クローズされたチャンネルに値を送信するとランタイムパニックになる。
 クローズされたチャンネルから受信した場合、そのチャネルの型の初期値が受信され続ける。
 チャネルからの受信の戻り値の内、２つ目はチャネルがクローズされていて、かつバッファが空の場合falseを表す。
 この値でチャネルがクローズされて受信するものが残っていないかを判断する。

 */

func main() {
	ch := make(chan int)

	go func() {
		count := 1
		for {
			time.Sleep(1 * time.Second)
			value := int(time.Now().Unix() )
			ch <- value
			count += 1

			if count > 4 {
				break
			}
		}
		close(ch)
		// ch <- 1 クローズされたチャンネルに値を送信するとランタイムパニックになる。

		fmt.Println("ch closed")
	}()

	for {
		time.Sleep(1 * time.Second)
		i, ok := <-ch
		fmt.Println(i, ok)
		if !ok {
			break
		}
	}
}


