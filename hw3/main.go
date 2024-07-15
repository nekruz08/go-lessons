package main

import "fmt"

func main() {

	PrintGreeting()
	DisplaySeparator()
	NotifyStart()
	result := GetWelcomeMessage()
	fmt.Println(result)
	result1 := GetPiValue()
	fmt.Println(result1)
	result2 := IsServerActive()
	fmt.Println(result2)
	PrintNumber(1000000)
	GreetUser("Muhammad ibn Musa al-Khwarizmi")
	ToggleLight(true)
	result3 := Add(1, 2)
	fmt.Println(result3)
	result4 := Concat("О", "Б")
	fmt.Println(result4)
	result5 := IsEven(7)
	fmt.Println(result5)
	result6 := adder(1, 2)
	fmt.Println(result6)
	result7 := concatenator("MO","DAR")
	fmt.Println(result7)
	

}

// 2.	В функции не передаются аргументы и ничего не возвращает функция
// Создайте функцию PrintGreeting, которая печатает "Hello, World!" на экран.
func PrintGreeting() {
	fmt.Println("Hello, World!")
}

// Создайте функцию DisplaySeparator, которая печатает строку из 20 символов - для разделения текста.
func DisplaySeparator() {
	fmt.Println("--------------------")
}

// Создайте функцию NotifyStart, которая выводит на экран сообщение "Program started".
func NotifyStart() {
	fmt.Println("Program started")
}

// В функции не передаются аргументы, но функция возвращает значение (значения)
// Создайте функцию GetWelcomeMessage, которая возвращает строку "Welcome!".
func GetWelcomeMessage() string {
	return "Welcome!"
}

// Создайте функцию GetPiValue, которая возвращает значение числа π с точностью до двух знаков после запятой (3.14).
func GetPiValue() float32 {
	return 3.14
}

// Создайте функцию IsServerActive, которая возвращает булево значение true.
func IsServerActive() bool {
	return true
}

// В функции передаются аргументы, но функция ничего не возвращает
// Cоздайте функцию PrintNumber, которая принимает целое число и выводит его на экран.
func PrintNumber(a int) {
	fmt.Println(a)
}

// Создайте функцию GreetUser, которая принимает строку с именем пользователя и выводит приветствие.
func GreetUser(s string) {
	fmt.Println(s)
}

// Создайте функцию ToggleLight, которая принимает булевое значение и выводит "Light on" или "Light off" в зависимости от значения аргумента.
func ToggleLight(b bool) {
	if b {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

// В функции передаются аргументы, но функция возвращает значение(значения)
// Создайте функцию Add, которая принимает два целых числа и возвращает их сумму.
func Add(a, b int) int {
	return a + b
}

// Создайте функцию Concat, которая принимает две строки и возвращает их конкатенацию.
func Concat(a, b string) string {
	return a + b
}

// Создайте функцию IsEven, которая принимает целое число и возвращает true, если число четное, и false в противном случае.
func IsEven(a int) bool {
	return a%2 == 0
}

// Функция - как переменная
// Создайте переменную adder, которая является функцией, принимающей два целых числа и возвращающей их сумму.
// Определение типа функции AdderFunc
type AdderFunc func(a, b int) int

var adder AdderFunc = func(a, b int) int {
	return a + b
}

//Создайте переменную concatenator, которая является функцией, принимающей две строки и возвращающей их конкатенацию.
type Concatenator func(a,b string) string

var concatenator Concatenator=func (a,b string) string {
	return a+b
}

// Создайте переменную isPositive, которая является функцией, принимающей целое число и возвращающей true, если число положительное.
