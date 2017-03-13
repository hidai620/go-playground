package main

import "fmt"

var p = func(s interface{}) {
	fmt.Printf("%#v\n", s)
}

/*
文字列型をfor rangeループで回した場合、
indexの値はbyte数分増加する。
 */
func main() {
	for i, s := range "アイウエオABCDE" {
		fmt.Printf("[%d] %d\n", i, s)
	}
	/*
		[0] 12450
		[3] 12452
		[6] 12454
		[9] 12456
		[12] 12458
		[15] 65
		[16] 66
		[17] 67
		[18] 68
		[19] 69
	 */
}
