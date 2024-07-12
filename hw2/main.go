package main

import (
	"fmt"
	"math"
)

var a int
var b int
var c int
var d int
var P float32 = 3.14
var R int

func main() {
	//1----------------------
	P := 4 * a
	fmt.Println(P)
	//2----------------------
	S := a ^ 2
	fmt.Println(S)
	//3----------------------
	S = a * b
	P = 2 * (a + b)
	fmt.Println(S)
	fmt.Println(P)
	//4----------------------
	L := P * d
	fmt.Println(L)
	//5----------------------
	V := a ^ 3
	S = 6 * a
	fmt.Println(V)
	fmt.Println(S)
	//6----------------------
	V = a * b * c
	S = 2 * (a*b + b*c + a*c)
	fmt.Println(V)
	fmt.Println(S)
	//7----------------------
	L = 2 * P * R
	S = P * R
	fmt.Println(L)
	fmt.Println(S)
	//8----------------------
	M := (a + b) / 2
	fmt.Println(M)
	//9----------------------
	G := math.Sqrt(float64(a) * float64(b))
	fmt.Println(G)
	//10----------------------
	if a&b > 0 {
		fmt.Println(a + b)
		fmt.Println(a - b)
		fmt.Println(a * b)
		fmt.Println(a ^ 2/b ^ 2)
	}
	//11----------------------
	// var L int
	// meters := L / 1000
	// fmt.Println(meters)
	//12----------------------
	// var M int
	// tons:=M/1000
	// fmt.Println(tons)
	//13----------------------
	// var B int
	// kilobytes := B / 1024
	// fmt.Println(kilobytes)
	//14----------------------
	// var A int
	// var B int
	// if A > B {
	// 	count := A / B
	// 	fmt.Println(count)
	// }
	//15----------------------
	var A int=16
	var B int=7
	if A > B {
		remainder := A % B
		fmt.Println(remainder)
	}

}
