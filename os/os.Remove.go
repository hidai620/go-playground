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

	err = os.Remove("./os/create_test")
	if err != nil {
		return
	}
}

