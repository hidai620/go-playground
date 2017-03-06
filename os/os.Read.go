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

	file, err := os.Open("./os/hello")
	if err != nil {
		return
	}

	bs := make([]byte, 128)
	i, err := file.Read(bs)
	if err != nil {
		return
	}

	fmt.Println(i)
	fmt.Println(string(bs))
}



