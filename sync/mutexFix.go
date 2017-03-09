package main

import (
	"fmt"
	"time"
	"sync"
)

type Box struct {
	a, b, c int
}


var box Box
var mutex *sync.Mutex

func updateWithMutex(i, count int) {
	mutex.Lock() // 一つのゴルーチンのみがロックを取得できる仕組み

	box.a = count
	time.Sleep(time.Millisecond)
	box.b = count
	time.Sleep(time.Millisecond)
	box.c = count
	time.Sleep(time.Millisecond)
	fmt.Printf("%i %#v \n",i, box)

	mutex.Unlock()
}

// レースコンディションでboxの中の値が一致しなくなる
// 非同期処理はWaitGroupで待ち受ける
// Addで同期する対象のスレッド数を登録し、
// その数だけ、Doneが呼び出されたらWaitを抜ける。
func main() {
	mutex = new(sync.Mutex)

	parallelism := 5

	wg := new(sync.WaitGroup)
	wg.Add(parallelism)

	for i := 0; i < parallelism; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				updateWithMutex(i, j)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
