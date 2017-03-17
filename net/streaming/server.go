package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"bufio"
	"os"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req*http.Request) {

		fmt.Println(req.Method, req.URL)
		printValues("HEADER", req.Header)

		file, err := ioutil.TempFile(os.TempDir(),"upload")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("uploaded file:", file.Name())

		writer := bufio.NewWriter(file)
		reader := bufio.NewReader(req.Body)

		reader.WriteTo(writer)
		writer.Flush()
		// レスポンスボディの作成
		res.Write([]byte("hello"))
	})

	http.ListenAndServe(":8080", nil)
}

func printValues(name string, values interface{}) {
	switch t := values.(type) {
	case http.Header:
		for k, v := range t {
			fmt.Printf("[%5s] %30s : %s \n", name, k, v)
		}
	case url.Values:
		for k, v := range t {
			fmt.Printf("[%5s] %30s : %s \n", name, k, v)
		}
	case map[string][]string:
		for k, v := range t {
			fmt.Printf("[%5s] %30s : %s \n", name, k, v)
		}
	}

}

//readBytes := 1024
//
//tmp := make([]byte, readBytes)
//for {
//	readed, err := reader.Read(tmp)
//	if err != nil {
//		fmt.Println("Read ERROR:", err)
//		if reader.Buffered() != 0 {
//			fmt.Println("Bufferd:", reader.Buffered())
//			writer.Flush()
//		}
//		return
//	}
//	fmt.Println("Readed:", readed)
//
//	if readed != 0 {
//		writer.Write(tmp)
//		writer.Flush()
//	} else {
//		// writer.Flush()
//		break
//	}
//}

