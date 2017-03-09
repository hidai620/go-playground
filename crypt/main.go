package main

import (
	"crypto/md5"
	"io"
	"fmt"
	"hash"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)


var p = fmt.Println

func showHexString(h hash.Hash) {
	io.WriteString(h, "ABCDE")

	// MD5ハッシュ値の生成
	bytes := h.Sum(nil)

	// 16進数の文字列を生成
	// s := fmt.Sprintf("%10s %x",reflect.TypeOf(h).String(), bytes)
	s := fmt.Sprintf("%10T %x",h , bytes)
	p(s)
}


func main() {
	showHexString(md5.New())
	showHexString(sha1.New())
	showHexString(sha256.New())
	showHexString(sha512.New())
}
