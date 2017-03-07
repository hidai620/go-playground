package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsDigit('9'))  // true
	fmt.Println(unicode.IsDigit('９')) // true
	fmt.Println(unicode.IsDigit('X'))  // false

	fmt.Println(unicode.IsLetter('X')) // true
	fmt.Println(unicode.IsLetter('あ')) // true
	fmt.Println(unicode.IsLetter('〒')) // false
	fmt.Println(unicode.IsLetter('9')) // false

	fmt.Println(unicode.IsSpace(' ')) // true
	fmt.Println(unicode.IsSpace('　')) // true
	fmt.Println(unicode.IsSpace('\t')) // true

	fmt.Println(unicode.IsSymbol('〒')) // true

	fmt.Println("-------------------------------") // true


	fmt.Println(containsSymbol("175-0091"))
	fmt.Println(containsSymbol("〒175-0091"))
	fmt.Println(containsSymbol("175-0091〒"))



}

func containsSymbol(s string) (bool) {
	for _,v := range s {
		isSymbol := unicode.IsSymbol(v)
		if isSymbol {
			return true
		}
	}
	return false
}
