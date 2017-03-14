package main

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

var p = fmt.Println

type accountId int

func NewAccountId(id int) (accountId, error) {
	var a accountId

	a = accountId(id)
	if err := a.Validate(); err != nil {
		return a, err
	}

	return a, nil
}


// メソッドはtypeで指定したエイリアスにも設定できる。
func (a accountId) String() string {
	return strconv.FormatInt(int64(a), 10)
}

func (a accountId) Validate() error {
	s := a.String()
	if len(s) != 11 {
		return errors.New("11桁でなければいけません")
	}

	if ! strings.HasPrefix(s, "71") {
		return errors.New("71で始まらなければいけません")
	}

	return nil

}

func main() {

	userId, err := NewAccountId(71000000000)
	if err != nil {
		println(err)
	}
	p(userId)

	userId2, err := NewAccountId(70000000000)
	if err != nil {
		fmt.Println(userId2, err.Error())
	}
	p(userId2)

	// メソッドを関数として取得
	f := accountId.String
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", f(userId))// 関数として実行

	f2 := userId.String
	fmt.Printf("%#v\n", f2())// 関数として実行



}
