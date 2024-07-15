package main

import "fmt"

func IsEven(number int) bool {
    return number%2 == 0
}

func main() {
    // Примеры вызова функции IsEven
    fmt.Println(IsEven(4))   // Вывод: true (4 является четным числом)
    fmt.Println(IsEven(7))   // Вывод: false (7 не является четным числом)
    fmt.Println(IsEven(-2))  // Вывод: true (-2 является четным числом)
}
