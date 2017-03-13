package main

import "fmt"

var p = func(s ...interface{}) {
	fmt.Printf("%d\n", s)
}

// アサーションは実行エラーを起こす
// interface型意外はアサーションを実行できない。コンパイルエラーになる
// switch文でアサーションした型による分岐ができる
// switch文でアサーションした型による分岐で値を利用することもできる。
// case 節で型名を列挙すると、値を取り出してもinterface型として取得してしまう。
func main() {

	var n interface{} = 3
	p(n.(int))
	_, ok := n.(int64) // 実行時エラーになる
	if !ok {
		fmt.Println("assertion failed")
	}

	// インターフェース型意外はアサーションできない。
	//var m int = 3
	//result, ok = m.(int64) // 実行時エラーになる


	// switch文でアサーションした型による分岐ができる
	var m interface{} = 10
	switch m.(type) {
	case int:
		fmt.Println("m is int")
	case float32:
		fmt.Println("m is float32")
	case string:
		fmt.Println("m is string")
	}

	var x interface{} = 10
	switch v := x.(type) { // 値を得ることもできる。
	case int:
		fmt.Println("m is int",v)
	case float32:
		fmt.Println("m is float32")
	case string:
		fmt.Println("m is string")
	}

}
