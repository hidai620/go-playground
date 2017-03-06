package main

import (
	"log"
	"os"
	"fmt"
)

func main() {

	logFile, err := os.Create("./log/test2.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	logger := log.New(logFile, "", log.Ldate | log.Ltime | log.Llongfile)


	log.Println("output std err")


	logger.Println("output logfile")
}
