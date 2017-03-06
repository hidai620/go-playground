package main

import "fmt"

func main() {
	fmt.Printf("[%d]\n", 123)   // [123]
	fmt.Printf("[%5d]\n", 123)  // [  123]
	fmt.Printf("[%-5d]\n", 123) // [123  ]
	fmt.Printf("[%05d]\n", 123) // [00123]

	fmt.Printf("[%f]\n", 123.456) // [123.456000]
	fmt.Printf("[%F]\n", 123.456) // [123.456000]
	fmt.Printf("[%.2f]\n", 123.456) // [123.46] // 小数点３桁を四捨五入
	fmt.Printf("[%8.2f]\n", 123.456) // [123.46] // 小数点３桁を四捨五入、全体を8桁で左詰め


	fmt.Printf("[%s]\n", "Go") // [Go]
	fmt.Printf("[%q]\n", "Go") // ["Go"]
	fmt.Printf("[%t]\n", true) // ["true"]
	fmt.Printf("[%T]\n", true) // [bool]

	fmt.Printf("[%v]\n", []string{"a", "b", "c"})

	fmt.Printf("[%v]\n", []string{"a", "b", "c"})
	fmt.Printf("[%v]\n", User{id: 1, name:"Taro"}) //[{1 Taro}]

	fmt.Printf("[%+v]\n", User{id: 1, name:"Taro"}) // [{id:1 name:Taro}]

	fmt.Printf("[%#v]\n", User{id: 1, name:"Taro"}) // [main.User{id:1, name:"Taro"}]
}

type User struct {
	id int
	name string
}
