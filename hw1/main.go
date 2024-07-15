package main

import (
	"fmt"
	"math"
)

func main() {
	// B E F O R E
	//int
	var a int
	var b int8
	var c int16
	var d int32
	var e int64
	//float
	var f float32
	var g float64
	//bool
	var h bool
	//rune
	var i rune
	//byte
	var j byte
	//string
	var k string
	//uint
	var l uint32
	var m uint64
	//complex
	var n complex64
	var o complex128

	fmt.Printf("* * * * *BEFORE* * * * *\n")
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", g, g)
	fmt.Printf("%T %v\n", h, h)
	fmt.Printf("%s %T %v\n", "rune", i, i)
	fmt.Printf("%s %T %v\n", "byte", j, j)
	fmt.Printf("%T %v\n", k, k)
	fmt.Printf("%T %v\n", l, l)
	fmt.Printf("%T %v\n", m, m)
	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", o, o)

	// A F T E R
	//int
	a = 9223372036854775807
	b = 127
	c = 32767
	d = 2147483647
	e = 9223372036854775807
	//float
	f = 3.4e+38  //3.4*10^38
	g = 1.4e+308 //1.8*10^308
	//bool
	h = true
	//rune // UTF 8
	i = 65
	//byte // ASCII
	j = 65
	//string
	k = "hello world"
	//uint
	l = math.MaxUint32
	m = math.MaxUint64
	//complex
	var maxFloat32 = float32(math.MaxFloat32)
	n = complex(maxFloat32, maxFloat32)
	var maxFloat64 = float64(math.MaxFloat64)
	o = complex(maxFloat64, maxFloat64)

	fmt.Printf("\n")
	fmt.Printf("* * * * *AFTER* * * * *\n")
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", g, g)
	fmt.Printf("%T %v\n", h, h)
	fmt.Printf("%s %T %v\n", "rune", i, string(i))
	fmt.Printf("%s %T %v\n", "byte", j, string(j))
	fmt.Printf("%T %v\n", k, k)
	fmt.Printf("%T %v\n", l, l)
	fmt.Printf("%T %v\n", m, m)
	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", o, o)

}
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b int
// 	a=4
// 	b=2
// 	fmt.Printf("%v+%v=%v\n",a,b,a+b)
// 	fmt.Printf("%v-%v=%v\n",a,b,a-b)
// 	fmt.Printf("%v*%v=%v\n",a,b,a*b)
// 	fmt.Printf("%v/%v=%v\n",a,b,a/b)
// }