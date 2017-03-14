package main

import "fmt"

type (
	Id int
	Age int
)

var p = fmt.Println

type callback func (i int) int

/*
typeの機能

エイリアスの定義
基本型に別名をつけることができる。
別名で値を生成できる。

 */
func main() {

	var id Id = 10
	p(id)

	id2 := Id(20)
	p(id2)


	// intのエイリアスとして定義したId, Age間に互換性はない。
	id3 := Id(30)
	age := Age(10)
	// id = age // cannot use age (type Age) as type Id in assignment

	// 明示的に型変換すれば代入できる
	id3 = Id(age)
	p(id3)

	// エイリアスの元の型に対しても互換性がない。
	i := 10
	// i = id // cannot use age (type Age) as type Id in assignment

	// 代入するためには明示的な型変換が必要。
	i = int(age)
	p("age to int",i)




	// 引数の関数にtypeで定義した別名を使用できる。
	arrays := [][]int {
		{1,2,3},
		{4,5,6},
	}

	result := flatMap(arrays, func(i int) int {
		return i * i
	})
	p(result)
}

// 引数の関数にtypeで定義した別名を使用できる。
func flatMap(arrays [][]int, callback callback)[]int {
	ret := make([]int, 0, 1000)
	for _, inner := range arrays {
		for _, v := range inner {
			ret = append(ret, callback(v))
		}
	}
	return ret
}
