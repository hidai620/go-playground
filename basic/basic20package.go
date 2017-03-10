package main

import (
	. "github.com/hidai620/go-playground/basic/basic20"
	//. "fmt"
)

/*
ドットで全体をインポートするとき、メソッドが衝突するとエラーになる。
Println redeclared during import "fmt"
	previous declaration during import "github.com/hidai620/go-playground/basic/basic20"
複数のファイルで一つのパッケージを構成するとき、importは定義したファイル内のみで有効。
 */
func main() {
	Println("A")
	//Println(1)
}
