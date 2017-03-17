package main

import (
	"fmt"
	"os"
)

// ファイルにフォーマットされた値を出力する。
func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Create("./fmt/write_test")
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "[%s]", "Hello")
}
