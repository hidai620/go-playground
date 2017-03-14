package main

import "fmt"

type T0 struct {
	Name0 string
	Id int
}
type T1 struct {
	Name1 string
	Id int
	T0
}
type T2 struct {
	Name2 string
	Id int
	T1
}


var p = func(x interface{}) {
	fmt.Printf("%#v\n", x)
}

/*
構造体を埋め込むと、
埋め込んだ構造体のプロパティ名がユニークになる場合に限り、
プロパティ名に直接アクセスできる。
ユニークにならないものは、構造体のプロパティ名を指定する必要がある。

埋め込みによって、オブジェクト指向の合成のようなことができる。
 */
func main() {

	var t2 = T2{
		Name2: "T2 Name",
		T1: T1 {
			Name1:"T1 Name",
			T0 : T0 {
				Name0 : "T0 Name",
			},
		},
	}
	p(t2)
	p(t2.Name0)
	p(t2.Name1)
	p(t2.Name2)
	p(t2.Id)
	p(t2.T1.Id)
	p(t2.T1.T0.Id)
}
