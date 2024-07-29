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
// package main

// import "fmt"

// func main() {
//     First(1)
// }

// func First(a int)  {
//     fmt.Println("Hello Wrold")
// }
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
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var input string
// 	fmt.Print("Enter the score: ")
// 	_, err := fmt.Scanln(&input)
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	score, err := strconv.Atoi(input)
// 	if err != nil {
// 		fmt.Println("Error converting input to integer:", err)
// 		return
// 	}

// 	grade := GetGrade(score)
// 	fmt.Printf("Grade for score %d is %s\n", score, grade)
// }

// func GetGrade(score int) string {
// 	var grade string
// 	switch {
// 	case score >= 90:
// 		grade = "A"
// 	case score >= 80:
// 		grade = "B"
// 	case score >= 70:
// 		grade = "C"
// 	default:
// 		grade = "F"
// 	}
// 	return grade
// }

// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("hello")
// 		if i==5{
// 			break
// 		}
// 	}
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
//     // Example 1: Skipping odd numbers
//     fmt.Println("Example 1: Skipping odd numbers")
//     for i := 1; i <= 10; i++ {
//         if i%2 != 0 {
//             // Skip odd numbers
//             continue
//         }
//         fmt.Println(i)
//     }

//     fmt.Println()

//	    // Example 2: Using continue with nested loops
//	    fmt.Println("Example 2: Using continue with nested loops")
//	    for outer := 1; outer <= 3; outer++ {
//	        fmt.Printf("Outer loop iteration %d\n", outer)
//	        for inner := 1; inner <= 3; inner++ {
//	            if outer == 2 && inner == 2 {
//	                // Skip printing "2 2"
//	                continue
//	            }
//	            fmt.Printf("%d %d\n", outer, inner)
//	        }
//	    }
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
//     var a *int
//     fmt.Println(*a)
// }
// -------------------------------------------------------------------------------------------------

package main

import "fmt"

func main() {
    var ptr *int // объявляем указатель на int
    fmt.Println("Значение указателя ptr:", ptr)
    
    if ptr == nil {
        fmt.Println("Указатель ptr равен nil")
    } else {
        fmt.Println("Указатель ptr не равен nil")
    }
}