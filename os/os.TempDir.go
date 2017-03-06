package main

import (
	"os"
	"fmt"
)

func main() {
	dir := os.TempDir()
	fmt.Println(dir)
}