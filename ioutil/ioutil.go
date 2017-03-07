package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	bytes, err := ioutil.ReadFile("./io/file")
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range bytes {
		fmt.Print(string(v))
	}

	fmt.Println(string(bytes))
}
