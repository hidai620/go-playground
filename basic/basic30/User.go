package basic30

import "fmt"

var defaultLanguage string

func init() {
	defaultLanguage = "jp"
	fmt.Println("basic30.init executed")
}

type User struct {
	id int
	name string
	language string
}

func NewUser(id int, name string) *User{
	return &User {
		id: id,
		name: name,
		language: defaultLanguage,
	}
}

