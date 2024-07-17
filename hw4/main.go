// Задания для 4 – го урока (начальный уровень)
package main

import "fmt"

func main() {
	PrintGreeting()
}

// Создайте функцию PrintGreeting, которая печатает
// "Доброе утро!", если dayType равен "утро";
// "Добрый день!", если dayType равен "день"; и
// "Добрый вечер!", если dayType равен "вечер".
//  Переменную dayType вводить с консоли внутри функции.

func PrintGreeting() {
	var dayTypem string
	fmt.Scanln(&dayTypem)

	if dayTypem == "утро" {
		fmt.Println("Доброе утро!")
	}  
	if dayTypem == "день" {
		fmt.Println("Добрый день!")
	}  
	if dayTypem == "вечер" { 
		fmt.Println("Добрый вечер!")
	}
}
