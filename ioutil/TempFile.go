package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func removeFile(file *os.File) {
	if file != nil {
		err := os.Remove(file.Name())// fileのNameメソッドはフルパスが帰る
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	var file *os.File
	// deferに登録できる関数は引数なしの関数。defer内で値を参照する場合クロージャする必要がある。
	defer func() {
		removeFile(file)
	}()

	//fmt.Println(os.TempDir())
	file, err := ioutil.TempFile(os.TempDir(),"foo")
	if err != nil {
		fmt.Println(err)
		return
	}

	file.WriteString("Hello")
}
