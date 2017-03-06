package main

import (
	"log"
	"os"
	"fmt"
)

func main() {

	// stderrに出力
	log.Println("std error")

	// stdoutに出力
	log.SetOutput(os.Stdout)
	log.Println("std out")


	file, err := os.Create("./log/test.log")
	if err != nil {
		fmt.Println(err)
	}

	log.SetOutput(file)
	log.Println("ファイルに出力")
}
