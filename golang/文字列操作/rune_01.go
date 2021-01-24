package main

import "fmt"

func main() {
	str := "abscdfg"
	// 文字列にインデックスでアクセスすると、対象の文字のバイト(utf-8)を表示。
	fmt.Printf("%v %T\n", str[0], str[0]) // 97 uint8

	// runeでキャストすることで、code pointのエイリアスであるint32を取ることができる。
	fmt.Printf("%v %T\n", rune(str[0]), rune(str[0])) // 97 int32

	// さらにstringでキャストすると、エイリアスであるint32から対応する文字列を取ることができる。
	fmt.Printf("%v %T\n", string(rune(str[0])), string(rune(str[0]))) // a string
}
