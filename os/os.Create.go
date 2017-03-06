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

	file, err := os.Create("./os/create_test")
	if err != nil {
		return
	}

	info, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(info.Mode())
	fmt.Println(info.IsDir())
	fmt.Println(info.Name())
	fmt.Println(info.Size())
}

