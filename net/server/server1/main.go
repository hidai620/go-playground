package main

import (
	"fmt"
	"net/http"
)

//func RootHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Println(req.URL, "RootHandler called!")
//	w.Write([]byte("RootHandler hello"))
//}
func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, "RootHandler called!")
	w.Write([]byte("RootHandler hello"))
}

func AnotherHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, "AnotherHandler called!")
	w.Write([]byte("AnotherHandler hello"))
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, "handler.ServeHTTP called!")
	w.Write([]byte("handler.ServeHTTP hello"))
}

/*
 組み込みのhttpパッケージのサーバー機能を試す

 httpパッケージの関数でサーバーをセットアップするパターン

 # HTTPサーバーを構築する手順
 ## URLハンドラーを登録する。
 ハンドラーを登録する方法は3種類ある。

 ### htt.HandleFuncを使う
 引数に渡すことができるのは、

 URLごとに、リクエストとレスポンスをハンドリングするハンドラー。
 "/"で登録すると、他に特定のパスを登録したハンドラーは、全てをハンドリングするハンドラーになる。
 "/a"で登録すると、"/a"に完全一致するURLをリクエストされた時に起動されるハンドラーになる。


 ### htt.Handleを使う。
 HandleFuncと同じようにURLごとに、リクエストとレスポンスをハンドリングするハンドラを渡す。
 net.Handlerインターフェースを満たす構造体を定義して渡す。

 type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
 }
 HandleFuncに渡した関数と同じ。

 ### http.ListenAndServeのを使う。
 第２引数にHandlerインターフェースを満たす構造対を渡す。
 ここにHandlerを渡すと、他にURLごとにHandlerを登録していても、
 全てが無視され、全てのURLでこのHandlerが実行される。

 ## htt.ListenAndServeでサーバーを起動する。
 最低限ポートだけ指定すればいい。
 URLにリクエストすると、登録したハンドラーが呼び出される。

*/
func main() {
	// "/"で登録すると、他に特定のパスを登録したハンドラーは、全てをハンドリングするハンドラーになる。
	http.HandleFunc("/", RootHandler)

	// "/a"に完全一致するURLをリクエストされた時に起動される。
	http.HandleFunc("/a", AnotherHandler)
	http.Handle("/b", handler{})
	http.ListenAndServe(":8081", handler{})

	mux := http.NewServeMux()
	mux.
}
