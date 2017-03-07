package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	id        int       //　プライベートフィールドなのでJSON出力されない
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"` // JSON出力されない
}

func main() {
	// Marshal 構造体を文字列にする
	u := new(User)
	u.id = 1
	u.Name = "Tom"
	u.CreatedAt = time.Now()

	bytes, err := json.Marshal(u) // パブリックパラメーターのみJSONに出力する。
	if err != nil {
		fmt.Println(err)
		return
	}

	println(string(bytes))

}
