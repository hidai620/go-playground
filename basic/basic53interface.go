package main

import (
	"io"
	"fmt"
)

// インターフェースも合成できる
type ReadWriter interface {
	io.Reader
	io.Writer
}

type MSWord struct {

}

func (t MSWord) Read(bytes []byte)(int, error) {
	fmt.Println("Now reading")
	return 0, nil
}

func (t MSWord) Write(bytes []byte)(int, error) {
	fmt.Println("Now writing", string(bytes))
	return 0, nil
}

func main() {

	document := new(MSWord)
	createReport(document)
}

func createReport(r ReadWriter) {

	r.Read(make([]byte, 128))

	r.Write([]byte("ABCDE"))
}

// typeで定義しなくても、引数に直接書くこともできる。
func showId(id interface{ GetId() int}) {
	fmt.Println(id.GetId())
}
