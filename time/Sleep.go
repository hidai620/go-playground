package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	time.Sleep(5 * time.Second)

	fmt.Println(time.Now())
}