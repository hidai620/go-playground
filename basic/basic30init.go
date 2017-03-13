package main

import (
	"fmt"
	"github.com/hidai620/go-playground/basic/basic30"
)

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

/*
initはパッケージの初期化を目的にしている。
パッケージをインポートした際に、パッケージに定義されたinitが呼ばれる。
次にmainに定義されたinitが呼ばれる。
initには引数や戻り値を定義することはできない。(init function must have no arguments and return values)
initは同名の関数を複数定義でき、定義された順番に実行されるが、可読性やメンテナンス性を考えて普通は１つだけ定義する。
 */
func init() {
	p("init 1")
}
func init() {
	p("init 2")
}

func main() {
	p("main")
	user := basic30.NewUser(1, "Tom")
	p(user)
}
