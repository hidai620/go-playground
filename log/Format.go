package main

import "log"

func main() {


	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("hello") // 2017/03/06 17:29:42 Format.go:9: hello

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("hello") // 2017/03/06 17:31:03 /Users/N1407A003/go/src/github.com/hidai620/go-playground/log/Format.go:12: hello
}
