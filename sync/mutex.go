package main

import (
	"fmt"
	"time"
)

type Box struct {
	a, b, c int
}


var box Box

func update(i, count int) {
	box.a = count
	time.Sleep(time.Millisecond)
	box.b = count
	time.Sleep(time.Millisecond)
	box.c = count
	time.Sleep(time.Millisecond)
	fmt.Printf("%i %#v \n",i, box)
}

// レースコンディションでboxの中の値が一致しなくなる
func main() {

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				update(i, j)
			}
		}()
	}
	for {

	}
}
