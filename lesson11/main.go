package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// var str1 string="Hello, World!"
	// fmt.Println(str1)

	// str2:="Hello, World!"
	// fmt.Println(str2)

	// str3:=`	Hello, World!
	// Hello, World!`
	// fmt.Println(str3)

	// s1:="Hello, "
	// s2:="World"
	// s3:=s1+s2
	// fmt.Println(s3)

	// s := "Hello"
	// c := s[1]
	// fmt.Println(c)         //101 байт с кодом
	// fmt.Println(string(c)) // 'e' приведение байт кодоа в string

	// s:="Hello, World!"
	// sub:=s[7:12] // поучаем значение по этим индексам т.е. тсроку World  12 не вклюсительно
	// fmt.Println(sub)

	// s1:="abc"
	// s2:="def"
	// fmt.Println(s1==s2)
	// fmt.Println(s1<s2)

	// s:="Hello, World!"
	// fmt.Println(strings.Contains(s,"World"))
	// fmt.Println(strings.Index(s,"World"))

	// s := "Hello, 世界"
	// bytes := []byte(s)
	// fmt.Println(bytes)

	// rune := []rune(s)
	// fmt.Println(rune)

	s := "Hello, 世界"
	fmt.Println(len(s)) // 13

	fmt.Println(utf8.RuneCountInString(s)) // 9 (при условии что счет начинается с нуля)

	for i, r := range s {
		fmt.Printf("%d: %c\n", i, r)
	}
}
