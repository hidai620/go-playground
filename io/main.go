package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	bytes := make([]byte, 10)

	for {
		io.ReadAtLeast(os.Stdin, bytes, 5)
		fmt.Println(string(bytes))
	}


}
