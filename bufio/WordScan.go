package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main() {
	s := `ABC DEFG
HIJK LMN
OPQR STU
`
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
