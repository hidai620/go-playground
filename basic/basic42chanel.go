package main

import (
	"time"
	"fmt"
)

var p = fmt.Println

/*
select 構文は複数のchanelを扱うための構文。
case 式は上から順番にではなく、ランダムに実行される。
 */
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go send(ch1)
	go send(ch2)
	go send(ch3)


	for {
		select {
		case v := <-ch1:
			p("received from ch1", v)
		case v := <-ch2:
			p("received from ch2", v)
		case v := <-ch3:
			p("received from ch3", v)
		}
	}

}

func send(ch chan int) {
	for {
		time.Sleep(1 * time.Second)
		ch <- int(time.Now().Unix())
	}
}

