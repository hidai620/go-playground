package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id        int       //　プライベートフィールドなのでJSON出力されない
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"` // JSON出力されない
	Phone     Phone
}

type Phone struct {
	Number int
	Type string
}


func main() {
	printModel(unmarshalUser("{\"name\":\"Tom\"}"))
	printModel(unmarshalUser("{\"name\":\"Tom\"}"))
	printModel(unmarshalUser("{\"Name\":\"Bob\"}"))

	// エラー Idを文字列型で取り込もうとするとエラー
	printModel(unmarshalUser("{\"id\" : \"9\", \"Name\":\"Bob\"}"))

	printModel(unmarshalUser("{\"id\" : 10, \"Name\":\"Bob\"}"))

 	// エラー シングルクォーテーションはダメ
	printModel(unmarshalUser("{'id' : 11, 'Name': 'Joe' }"))

	// 構造体を含むJSONもアンマーシャルできる
	printModel(unmarshalUser(`{"id" : 12, "Name":"Jobs",  "Phone": {"Number": 123456789, "Type": "iPhone" }   }`))

}

func printModel(i interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", i)
}

func unmarshalUser(jsonString string) (user *User, err error) {
	user = new(User)
	err = json.Unmarshal([]byte(jsonString), user)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
