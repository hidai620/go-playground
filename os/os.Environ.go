package main

import (
	"os"
	"fmt"
)

func main() {
	for i, v := range os.Environ() {
		fmt.Println(i, v)
	}


	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)

	gopath1 := os.Getenv("GOPATH1")
	fmt.Println(gopath1)

	gopath2, ok := os.LookupEnv("GOPATH2")
	if ok {
		fmt.Println(gopath2)
	} else {
		fmt.Println("GOPATH2は見つかりませんでした。")
	}
}