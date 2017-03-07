package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".css"))
	fmt.Println(mime.TypeByExtension(".jpg"))
	fmt.Println(mime.TypeByExtension(".png"))
}