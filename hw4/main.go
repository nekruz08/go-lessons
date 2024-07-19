// Задания для 4 – го урока (начальный уровень)
// package main

// import "fmt"

// func main() {
// 	PrintGreeting()
// }

// // Создайте функцию PrintGreeting, которая печатает
// // "Доброе утро!", если dayType равен "утро";
// // "Добрый день!", если dayType равен "день"; и
// // "Добрый вечер!", если dayType равен "вечер".
// //  Переменную dayType вводить с консоли внутри функции.

// func PrintGreeting() {
// 	var dayTypem string
// 	fmt.Scanln(&dayTypem)

//		if dayTypem == "утро" {
//			fmt.Println("Доброе утро!")
//		}
//		if dayTypem == "день" {
//			fmt.Println("Добрый день!")
//		}
//		if dayTypem == "вечер" {
//			fmt.Println("Добрый вечер!")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию PrintWeather, которая печатает
// "Солнечно", если weatherType равен "солнечно";
// "Облачно", если weatherType равен "облачно"; и
// "Дождливо", если weatherType равен "дождливо".
// Переменную weatherType вводить с консоли внутри функции.
// package main

// import "fmt"

// func main() {
// 	PrintWeather()
// }
// func PrintWeather() {
// 	var weatherType string
// 	fmt.Scanln(&weatherType)
// 	switch weatherType {
// 	case "солнечно":
// 		{
// 			fmt.Println("Солнечно")
// 		}
// 	case "облачно":
// 		{
// 			fmt.Println("Облачно")
// 		}
// 	case "дождливо":
// 		{
// 			fmt.Println("Дождливо")
// 		}
// 	}
// }

// -------------------------------------------------------------------------------------------------
// Создайте функцию PrintTrafficLight, которая печатает
// "Стоп", если lightColor равен "красный";
// "Внимание", если lightColor равен "желтый"; и
// "Идите", если lightColor равен "зеленый".
// Переменную lightColor вводить с консоли внутри функции.
// package main

// import "fmt"

// func main() {
// 	lightColor()
// }
// func lightColor()  {
// 	var lightColor string
// 	fmt.Scanln(&lightColor)

//		switch lightColor{
//		case "красный":
//			fmt.Println("Стоп")
//		case "желтый":
//			fmt.Println("Внимание")
//		case "зеленый":
//			fmt.Println("Идите")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию GetGrade, которая возвращает оценку "A", "B" или "C"
// в зависимости от значения переменной score.
// Переменную scope вводить с консоли внутри функции.

// значения берите от 0 до 100
// >=90  A, между 75 и 90 B, меньше 75 с
// возвращать A,B или С2
// package main

// import "fmt"

// func main() {
// fmt.Println(GetGrade())
// }
// func GetGrade() string {
// 	var score int
// 	fmt.Scan(&score)

// 		var grade string
// 		switch {
// 		case score >= 90:
// 			grade = "A"
// 		case score >= 75 && score < 90:
// 			grade = "B"
// 		case score < 75:
// 			grade = "C"
// 		}
// 		return grade
// 	}
// -------------------------------------------------------------------------------------------------
// Создайте функцию GetDiscount,
// которая возвращает скидку
// "10%" - при amount > 1000,
// "5%" - при 500 < amount <=100 или
// "0%" - при amount <= 500
// в зависимости от суммы покупки amount.
// Переменную amount вводить с консоли внутри функции.
// package main

// import "fmt"

//	func main() {
//		fmt.Println(GetDiscount())
//	}
//
//	func GetDiscount() int {
//		var amount int
//		fmt.Scan(&amount)
//		switch{
//		case amount>1000:
//			return 10
//		case amount >=100:
//			return 5
//		default:
//			return 0
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию GetTemperatureDescription,
// которая возвращает
// "Холодно" (меньше 10) ,
// "Тепло" (с 10 до 25) или
// "Жарко" в зависимости от значения переменной temperature.
// Переменную temperature вводить с консоли внутри функции.
// package main

// import "fmt"

// func main() {
// 	GetTemperatureDescription()
// }
// func GetTemperatureDescription() {
// 	var temperature int
// 	fmt.Scan(&temperature)
// 	switch {
// 	case temperature < 10:
// 		fmt.Println("Холодно")
// 	case temperature >=10 && temperature <=25:
// 		fmt.Println("Тепло")
// 	default:
// 		fmt.Println("Жарко")
// 	}

// }
// -------------------------------------------------------------------------------------------------
// Создайте функцию CheckNumber, которая принимает
// целое число и печатает
// "Положительное",если число больше нуля;
// "Отрицательное", если меньше нуля; и
// "Ноль", если равно нулю.
// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Scan(&num)
// 	CheckNumber(num)
// }
// func CheckNumber(n int) {
// 	switch {
// 	case n > 0:
// 		fmt.Println("Положительное")
// 	case n < 0:
// 		fmt.Println("Отрицательное")
// 	case n == 0:
// 		fmt.Println("Ноль")

//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию CheckAge, которая принимает возраст и печатает
// "Совершеннолетний", если возраст 18 и старше;
// "Несовершеннолетний", если младше 18.
// package main

// import "fmt"

//	func main() {
//		var age int
//		fmt.Scan(&age)
//		CheckAge(age)
//	}
//
//	func CheckAge(age int) {
//		switch {
//		case age >= 18:
//			fmt.Println("Совершеннолетний")
//		case age < 18:
//			fmt.Println("Несовершеннолетний")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию CheckPassword, которая принимает строку пароля и печатает
// "Пароль верный", если пароль равен "1234"; и
// "Пароль неверный" в противном случае.
// package main

// import "fmt"

// func main() {
// var pass string
// fmt.Scanln(&pass)
// CheckPassword(pass)

// }

//	func CheckPassword(pass string) {
//		switch{
//		case pass=="1234":
//			fmt.Println("Пароль верный")
//		default:
//			fmt.Println("Пароль неверный")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию Add, которая принимает два целых числа и возвращает сумму их модулей.
// Используйте условные операторы для проверки значений чисел.
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(Add(a, b))
}
func Add(a, b int) int {
	var summ int
	if a >= 0 && b >= 0 {
		summ = a + b
	} else if a >= 0 && b <= 0 {
		summ = a + (b * -1)
	} else if a <= 0 && b >= 0 {
		summ = (a * -1) + b
	} else {
		summ = (a + b) * -1
	}
	return summ
}
