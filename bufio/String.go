package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main() {
	message := `aaa
bbb
ccc
`
	scanner := bufio.NewScanner(strings.NewReader(message))

	println(scanner)
	println(scanner)
	println(scanner)
	println(scanner)
}

func println(scanner *bufio.Scanner) {
	if scanner.Scan() {
		fmt.Println(scanner.Text())
	} else {
		fmt.Println("scan failed")
	}
}


