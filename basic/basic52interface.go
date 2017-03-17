package main

import "fmt"

type UUIDProvider interface{
	UUID() string
}

type User struct {
	Id int
	Name string
	uuid string
}

func (u *User) UUID() string {
	return u.uuid
}

type Admin struct {
	Id int
	Name string
	uuid string
	Role int
}

func (a *Admin) UUID() string {
	return a.uuid
}
// ポインタ型と値型の両方にメソッドを定義するとエラー
//func (a Admin) UUID() string {
//	return a.uuid
//}

/*
インターフェース型のスライスには、
その型のインターフェースを満たすインスタンスであれば、違う型の構造体を入れることができる。

取り出した時はインターフェース型になる。
元の型を知りたい場合、アサーションで取得する。
switch case では、構造体がポインタ型なのか、厳密にチェックされる。
 */
func main() {
	users := []UUIDProvider {
		&User{Id:1, Name:"Tom", uuid:"xxxxxxxx"},
		&Admin{Id:2, Name:"Joe", uuid:"yyyyyyy"},
	}

	for i, v := range users {
		fmt.Printf("%-5d %-10s %#v\n", i, v.UUID(), v)
		switch v.(type) {
		case *Admin:
			fmt.Printf("%s is admin(reference) \n", v.UUID())
		}
	}

}
