package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("./bufio/output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bw := bufio.NewWriter(file)
	defer file.Close()

	A:
	for {
		br := bufio.NewReader(os.Stdin)
		br = bufio.NewReaderSize(br, 4)

		bytes, isPrefix, err := br.ReadLine()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(isPrefix)
		s := string(bytes)

		switch s {
		case "exit":
			// break これだとswitchをブレイクするので無限ループは抜けられない
			// return  // or
			break A  // or
		default :
			fmt.Println(s)
			bw.WriteString(s + "\n")
			bw.Flush()
		}
	}
}
