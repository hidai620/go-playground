package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	mimeType := "application/octet-stream"
	res, err := http.Post("http://localhost:8080", mimeType, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	 defer res.Body.Close()

	byte := make([]byte, 128)
	_, err = res.Body.Read(byte)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byte))
}
