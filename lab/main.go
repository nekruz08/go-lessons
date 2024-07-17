// package main

// import "fmt"

// func IsEven(number int) bool {
//     return number%2 == 0
// }

// func main() {
//     // Примеры вызова функции IsEven
//     fmt.Println(IsEven(4))   // Вывод: true (4 является четным числом)
//     fmt.Println(IsEven(7))   // Вывод: false (7 не является четным числом)
//     fmt.Println(IsEven(-2))  // Вывод: true (-2 является четным числом)
// }
//-------------------------------------------------------------------------------------------------
// package main

// func main() {
// 	// fmt.Println()
// 	Calculate(1, 2, func(a, b int) int { return a + b })

// }

// // Функция - как аргумент функции
// // Создайте функцию Calculate, которая принимает два целых числа и функцию для их обработки.
// // Примените переданную функцию к этим числам и верните результат.
//
//	func Calculate(a, b int, oldFunc func(a, b int) int) int {
//		return oldFunc(a, b)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
//     // тут мы вызоваем нашу функию которую определили за пределами main
//     newfunc()
// }

// // эту фунуцию можем вызовать в main

// 	func newfunc()  {
// 	    fmt.Println("Hard Wordk")
// 	}

// -------------------------------------------------------------------------------------------------
// Создайте функцию Execute, которая принимает булевое значение и функцию,
// которая принимает булевое значение и ничего не возвращает.
// Выполните переданную функцию с переданным булевым значением.

// package main

// import "fmt"

// func main() {

// 	Execute(true, func(d bool) { fmt.Println(d) })
// }

// func Execute(a bool, boolFunc func(d bool)) {
// 	boolFunc(true)
// }

// -------------------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
    First(1)
}

func First(a int)  {
    fmt.Println("Hello Wrold")
}
// -------------------------------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	result:=Multiplier(80)
// 	result()
// }
// func Multiplier(factor int)( kfc func () ){
// return kfc(factor int){
	
// }

// }
