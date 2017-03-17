package main

import (
	"fmt"
)

type User struct {
	Name string
}


type MasterUser User


type BillingUser User


type AnotherUser struct {
	Name string
}

type SuperUser struct {
	Id string
	Name string
}

func main() {
	//s := "hello"
	//bytes := []byte(s)
	//for _, v := range bytes {
	//	fmt.Println(v)
	//}
	//println(bytes)

	ss := string([]byte{104, 101, 108, 108, 111})
	fmt.Println(ss)
	//fmt.Println(strconv.FormatBool(true))

	user := User{Name:"Tom"}
	fmt.Printf("User                       %#v\n", user)
	fmt.Printf("User        -> MasterUser  %#v\n", MasterUser(user))
	fmt.Printf("User        -> BillingUser %#v\n", BillingUser(user))


	billing := BillingUser{Name:"Joe"}
	fmt.Printf("BillingUser -> MasterUser  %#v\n", MasterUser(billing))
	fmt.Printf("BillingUser -> User        %#v\n", User(billing))


	// type でのエイリアス関係にない構造体も型変換できる
	anotherUser := AnotherUser{Name: "Bob"}
	fmt.Printf("AnotherUser                 %#v\n", anotherUser)
	fmt.Printf("AnotherUser  -> MasterUser  %#v\n", MasterUser(anotherUser))
	fmt.Printf("AnotherUser  -> BillingUser %#v\n", BillingUser(anotherUser))

}
