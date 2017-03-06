package main

import (
	"os"
	"fmt"
)

func main() {
	var err error
	defer func () {
		if err != nil {
			fmt.Println(err)
		}
	}()

	// ReadAtは読み込む対象の長さ以上のbyte配列を引数に指定すると必ずEOFのエラーになってしまう。

	file, err := os.Open("./os/hello")
	bs := make([]byte, 10)
	_, err = file.ReadAt(bs, 2)
	if err != nil {
		return
	}
	fmt.Println("4byte目から読み込み:", string(bs))
}



