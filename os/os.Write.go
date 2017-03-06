package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	file, err := os.OpenFile("./playground/hello", os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()


	i, err := file.WriteString("test test")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(i)


}



