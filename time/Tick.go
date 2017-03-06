package main

import (
	"time"
	"fmt"
)

func main() {
	// ２秒感覚で時間を受け取るチャンネルの生成
	ch := time.Tick(2 * time.Second)

	// 無限ループで２秒感覚で時刻を表示し続ける
	for {
		t := <- ch
		fmt.Println(t)
	}
}
