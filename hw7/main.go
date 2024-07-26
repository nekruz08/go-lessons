// Проверка температуры
// Определите тип Temperature на основе float64.
// Напишите функцию, которая принимает температуру и возвращает сообщение о том,
// является ли она ниже, выше или равной нулю.
package main

import "fmt"

type Temperature float64

func main() {
	fmt.Println(a(0))
}
func a(t Temperature) string {
	if t < 0{
		return "ниже нуля"
	}else if t >0{
		return "выше нуля"
	}else{
		return "равно нулю"
	}
}