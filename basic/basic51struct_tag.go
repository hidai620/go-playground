package main

import (
	"reflect"
	"fmt"
)



/*
タグはRAW文字列リテラルで描かれる傾向がある。
reflectパッケージで宣言されたタグを参照できる。
jsonパッケージはタグの情報を使って構造体からJSON文字列を作成する。
タグは文字列なので、コンパイル時に間違いを発見できない。
 */
type User struct {
	Id int `ID`
	Name string `名前`
}

func main() {

	u := User{Id:1, Name:"Tom"}

	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%-10s %-10s %-10s\n", field.Name, field.Type, field.Tag)
	}
}
