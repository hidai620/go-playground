package main

import (
	"time"
	"fmt"
)

/*
 # チャネル
 ## 用途
 ゴルーチン間で値を受け渡すためのキュー。
 デフォルトでチャネル内に格納できる値の数は１つ。

 # バッファ
 チャネルを生成するときにバッファサイズを指定できる。
 バッファを指定しない場合、バッファサイズがゼロのチャネルができ、チャネル内に格納できる値の数は１つ。

 バッファはあくまで、チャネルの追加の送受信領域であって、チャネル全体のキューのサイズを示すものではない。
 バッファがゼロのチャネルは送受信できる領域が1つ。（デフォルト領域:1 + バッファ:0）
 ch := make(chan int64) // 容量1

 バッファが1のチャネルは送受信できる領域が2つ（デフォルト領域:1 + バッファ:1)
 ch := make(chan int64, 1) // 容量2

 ## 送信側への影響
 送信側は、チャネルに値を送信したあと、
 バッファを含め、チャネル全体の格納領域に空きがある場合、送信後、次の処理に移れる。
 空きがない場合、送信処理の行で、受信されるまで待つことになる。

 ### 考えられる用途
 - バッファーの数を制限して、チャネル内のデータでのメモリの占有量を制限する。
 - チャネルへ送信する処理の方が受信側の処理より早い場合、チャネルの容量を調整することで、送信側をブロックしない。

 ## チャネル内のデータの有無
 len関数をチャネル型の変数に適用した場合、戻り値として得られるのは、バッファ内の値の件数。
 チャネル内全体の件数ではない。
 チャネル内全体の件数は、lenで取得した件数が1以上であれば、
 チャネルのデフォルトで格納できる1件 + lenで取得した件数で計算できる。
 lenが0を返す場合、かつチャネルのバッファがゼロの場合、
 チャネルに値が入っているかどうかlenの結果からでは判断がつかない。
 */

func main() {
	ch := make(chan int64)

	//go sender(ch)
	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			var value = time.Now().Unix()
			fmt.Printf("sender: send before %d\n", value)
			ch <- value  // バッファに空きがない状態では送ったあと、受信されるまで次の処理に進まない。
			fmt.Printf("sender: send after %d\n")
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("receiver: current buffering size %d\n", len(ch))
		// fmt.Printf("received %d\n", <-ch)
	}
}



func sender(ch chan int64) {
	for {
		time.Sleep(1 * time.Millisecond)
		var value = time.Now().Unix()
		fmt.Printf("sender: send before %d\n", value)
		ch <- value  // バッファに空きがない状態では送ったあと、受信されるまで次の処理に進まない。
		fmt.Printf("sender: send after %d\n")
	}
}
