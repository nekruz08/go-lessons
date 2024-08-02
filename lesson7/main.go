// code form lecture: readme.md file
// package main

// import "fmt"
// type Age int

// 	func main() {
// 		var myAge Age = 25
// 		var apple int = 25
// 		fmt.Println(myAge+Age(apple))
// 	}

// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Adder func(int, int) int

// func add(a int, b int) int {
// 	return a + b
// }

//	func main() {
//		var myAdder Adder = add
//		result := myAdder(2, 3)
//		fmt.Println(result)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Adder func(int, int) int

// func add(a int, b int) int {
// 	return a + b
// }
// func main() {
// 	var myAdder Adder = add
// 	result := myAdder(2, 3)
// 	fmt.Println(result)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type MyInt int

// func main() {
// 	var x MyInt = 10
// 	fmt.Println(x)
// 	fmt.Println(x + 1) // c литералом работает

// 	var y int = 1
// 	fmt.Println(x + y) // с другим типом не работает, например int
// }

// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// type MyInt=int

//	func main() {
//		var x MyInt=10
//		fmt.Println(x+2)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age int
// }

// func main() {
// 	p:=Person{Name: "Alice",Age: 30}
// 	fmt.Println(p)
// 	fmt.Printf("%+v\n",p)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// type Person struct{
// 	Name string
// 	Age int
// }
// func main() {
// 	p:=Person{Name: "Alice",Age: 30}
// 	fmt.Printf("%+v\n",p)
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// type Person struct{
// 	Name string
// 	Age int
// }
// func main() {
// 	p:=Person{"Alice",30}
// 	fmt.Println(p)
// 	fmt.Printf("%+v\n",p)
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// type Person struct{
// 	Name string
// 	Age int
// }
// func main() {
// 	p:=new(Person)
// 	p.Age=100
// 	p.Name="Nekruz"
// 	fmt.Printf("%+v\n",*p)
// 	fmt.Printf("%+v\n",p)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	p := &Person{Name: "Salom", Age: 22}
// 	fmt.Printf("%+v\n", *p)
// 	fmt.Printf("%+v\n", p)
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		p:=Person{Name: "Alice",Age: 30}
//		fmt.Println(p.Name)
//		fmt.Println(p.Age)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// type Person struct{
// 	Name string
// 	Age int
// }

// func main() {
// p:=&Person{Name:"Alice", Age: 30}
// p.Age=31
// fmt.Println(p)
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// type Address struct {
// 	City  string
// 	State string
// }
// type Person struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }
// func main() {
// 	p:=Person{
// 		Name: "Alice",
// 		Age: 30,
// 		Address: Address{
// 			City: "New York",
// 			State: "NY",
// 		},
// 	}
// 	fmt.Println(p)
// 	fmt.Printf("%+v\n",p)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// type Node struct{
// 	Value int
// 	Next *Node
// }
// func main() {
// 	n1:=&Node{Value: 1}
// 	n2:=&Node{Value: 2}
// 	n3:=&Node{Value: 3}

// 	n1.Next=n2
// 	n2.Next=n3

//		// Печатаем значениие узлов
//		current:=n1
//		for current!=nil{
//			fmt.Println(current.Value)
//			current=current.Next
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		person:=struct{Name string;Age int}{Name: "Alica",Age: 12}
//		fmt.Println(person)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	person:=struct{
// 		Name string
// 		Age int
// 		}{
// 			Name: "Alica",
// 			Age: 32,
// 		}
// 		fmt.Println(person)
// }

// problem solving

// -------------------------------------------------------------------------------------------------
// Определение возраста совершеннолетия
// Определите тип Age на основе int.
// Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
// package main

// import "fmt"
// type Age int

// 	func a(age int) string {
// 		if age >=18{
// 			return "совершеннолетний"
// 		}
// 		return "несовершеннолетний"
// 	}

// 	func main() {
// 		fmt.Println(a(13))

// 	}

// -------------------------------------------------------------------------------------------------
// Проверка на четность
// Определите тип Number на основе int.
// Напишите функцию, которая принимает число и возвращает сообщение о том,
// является ли оно четным или нечетным.
// package main

// import "fmt"
// type Number int
// func isEven(number Number) string {
// 	if number%2 == 0{
// 		return "четный"
// 	}
// 	return "нечетный"
// }
// func main() {
// 	fmt.Println(isEven(6))
// }

// -------------------------------------------------------------------------------------------------
// Проверка допустимого диапазона
// Определите тип Score на основе int. Напишите функцию,
// которая принимает оценку и возвращает сообщение,
// находится ли она в допустимом диапазоне от 0 до 100
// package main

// import "fmt"

// type Score int

// 	func isValidScore(score Score) string {
// 		if score >= 0 && score <=100 {
// 			return "Valid Score"
// 		}
// 		return "Invalid Score"
// 	}

// 	func main() {
// 		fmt.Println(isValidScore(1))
// 	}
// -------------------------------------------------------------------------------------------------
// Арифметические операции
// Определите тип функции Operation, которая принимает два int и возвращает int.
// Напишите функции для сложения,
// вычитания и умножения. Используйте тип Operation для вызова этих функций
// package main

// import "fmt"
// type Operation func (a,b int) int
// func adding(a,b int) int {
// 	return a+b
// }
// func subtract(a,b int) int {
// 	return a-b
// }
// func multiply(a,b int) int  {
// 	return a*b
// }
// func main() {
// 	var a Operation=adding
// 	fmt.Println(a(1,1))

// 	var b Operation=subtract
// 	fmt.Println(b(4,2))

// 	var c Operation=multiply
// 	fmt.Println(c(2,2))
// }
// -------------------------------------------------------------------------------------------------
// Сравнение чисел
// Определите тип функции Comparator, которая принимает два int и возвращает bool.
// Напишите функции для проверки
// равенства и сравнения больше/меньше. Используйте тип Comparator для вызова этих функций.
// package main

// import "fmt"

// type Comparator func(a, b int) bool

// func checkEquality(a, b int) bool {
// 	return a == b
// }

// func greater(a,b int) bool {
// 	return a>b
// }
// func less(a,b int) bool {
// 	return a<b
// }

// func main() {
// var a Comparator=checkEquality
// fmt.Println(a(2,2)) //true

// var b Comparator=greater
// fmt.Println(b(2,3)) //false

// var c Comparator=less
// fmt.Println(c(3,4)) // true
// }
// -------------------------------------------------------------------------------------------------
// Применение функции к числу
// Определите тип функции UnaryOperation, которая принимает int и возвращает int.
// Напишите функции для удвоения и утроения числа. Используйте тип UnaryOperation для вызова этих функций.
// package main

// import "fmt"
// type UnaryOperation func (int) int
// func duble(n int) int {
// 	return n*2
// }
// func triple(n int) int {
// 	return n*3
// }
// func main() {
// 	var a UnaryOperation = duble
// 	fmt.Println(a(2)) // 2*2=4

// 	var b UnaryOperation = triple
// 	fmt.Println(b(2)) // 2*2*2=8

// }
// -------------------------------------------------------------------------------------------------
// Обратный отсчет
// Создайте псевдоним для типа int под названием Count.
// Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля.
// package main

// import "fmt"
// type Count=int

// 	func countdown(c Count) {
// 		for i := c; i >=0; i-- {
// 			fmt.Println(i)
// 		}
// 	}

//	func main() {
//		countdown(10)
//	}
//
// -------------------------------------------------------------------------------------------------
// Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature.
// Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).
// package main

// import "fmt"

// type Temperature = float64

// func a(t Temperature) string {
// 	if t < 0 {
// 		return "Below zero"
// 	} else if t > 0 {
// 		return "Above zero"
// 	}
// 	return "Zero"
// }

// func main() {
// 	fmt.Println(a(-1))
// }
// -------------------------------------------------------------------------------------------------
// Применение скидки
// Создайте псевдоним для типа float64 под названием Price.
// Напишите функцию, которая принимает Price и возвращает новую цену с 10% скидкой.
// package main

// import "fmt"
// type Price = float64
// func a(p Price) float64 {
// 	// return (p-float64(p/10))
// 	// return (p/100)*90
// 	return p*0.9
// }
// func main() {
// 	fmt.Println(a(100))

// }

// -------------------------------------------------------------------------------------------------
// Информация о пользователе
// Создайте структуру User с полями Name (строка) и Age (целое число).
// Напишите функцию, которая принимает User и выводит информацию о нем.
// package main

// import "fmt"

// 	type User struct{
// 		Name string
// 		Age int
// 	}

// 	func printUser(u User) {
// 		// fmt.Printf("%+v\n",s)
// 		fmt.Printf("Name: %s, Age: %d\n",u.Name,u.Age)

// 	}

// 	func main() {
// 		b:=User{
// 			Name: "Abdurahmon",
// 			Age: 32,
// 		}
// 		printUser(b)
// 		printUser(User{"Alica",32})
// 	}
//-------------------------------------------------------------------------------------------------
// Продукт и его стоимость
// Создайте структуру Product с полями Name (строка) и Price (тип Price).
// Напишите функцию, которая принимает Product и возвращает сообщение о его стоимости.
// package main

// import "fmt"

// type Price float64

// 	type Product struct {
// 		Name  string
// 		Price Price
// 	}

// 	func printProduct(p Product)  {
// 		fmt.Printf("Product: %s, Price: %.2f\n",p.Name,p.Price)
// 	}

// 	func main() {
// 		var price = 12.5
// 		product:=Product{
// 			Name: "Каймок",
// 			Price: Price(price),
// 		}
// 		printProduct(product)
// 	}
//-------------------------------------------------------------------------------------------------
// Информация о книге
// Создайте структуру Book с полями Title (строка), Author (строка) и Pages (целое число).
// Напишите функцию,которая принимает Book и выводит информацию о книге.
// package main

// import "fmt"

// 	type Book struct{
// 		Title string
// 		Author string
// 		Pages int
// 	}

// 	func printBook(b Book) {
// 		fmt.Printf("%+v\n",b)
// 		fmt.Printf("Title: %s, Author: %s, Pages: %d\n",b.Title,b.Author,b.Pages)
// 	}

//	func main() {
//		var book Book=Book{
//			Title: "Who he was",
//			Author: "Nekruz",
//			Pages: 100,
//		}
//		printBook(book)
//	}
//
// -------------------------------------------------------------------------------------------------
// Изменение данных через указатель
// Создайте структуру Person с полями Name и Age.
// Напишите функцию, которая принимает указатель на Person
// и обновляет его возраст. Выведите информацию до и после обновления.
// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func a(p *Person,newAge int) {
// 	p.Age = newAge
// }
// func main() {
// 	p := Person{
// 		Name: "Nekruz",
// 		Age:  30,
// 	}
// 	fmt.Printf("%+v\n", p)
// 	a(&p,32)
// 	fmt.Printf("%+v\n", p)

// }

// -------------------------------------------------------------------------------------------------
// Создание и изменение структуры через указатель
// Создайте структуру Rectangle с полями Width и Height.
// Напишите функцию, которая принимает указатель на Rectangle,
// вычисляет и обновляет его площадь, а затем выводит обновленную площадь.
// package main

// import "fmt"

// 	type Rectangle struct{
// 		Width	int
// 		Height int
// 	}

// 	func updateArea(r *Rectangle)  {
// 		area:=r.Width*r.Height
// 		fmt.Printf("Area: %d\n",area)
// 	}

// 	func main() {
// 		updateArea(&Rectangle{2,2})
// 	}
// -------------------------------------------------------------------------------------------------
// Сравнение двух структур через указатели
// Создайте структуру Coordinate с полями X и Y. Напишите функцию, которая принимает указатели на две
// Coordinate и возвращает сообщение о том, равны ли они или нет.
// package main

// import "fmt"

// type Coordinate struct {
// 	X int
// 	Y int
// }

// func CompareCoordinates(a,b *Coordinate) string {
// 	if a.X==b.X&&a.Y==b.Y{
// 		return "Coordinates are equal"
// 	}
// 	return "Coordinates are not equal"
// }

// func main() {

// 	a := Coordinate{X: 1, Y: 2 }
// 	b := Coordinate{X: 1, Y: 2}

// 		fmt.Println(CompareCoordinates(&a,&b))
// 	}

// -------------------------------------------------------------------------------------------------
// Программа для управления студентом
// Создайте структуры Student и Address.
// Student должен содержать поля Name, Age и Address
// (вложенная структура Address).
// Напишите функцию, которая выводит информацию о студенте, включая адрес.
// package main

// import "fmt"

// type Student struct {
// 	Name   string
// 	Age    int
// 	Adress Adress
// }
// type Adress struct {
// 	Street string
// 	City	string
// }

// func printStudent(student Student)  {
// 	fmt.Printf("Name: %s, Age: %d\n",student.Name,student.Age)
// 	fmt.Printf("Adress: %s, %s\n",student.Adress.Street,student.Adress.City)
// }

//	func main() {
//		printStudent(Student{Name: "Nabot",Age: 21,Adress: Adress{Street: "Sino",City: "Dushanbe"}})
//	}
//
// -------------------------------------------------------------------------------------------------
// Работа с сотрудниками
// Создайте структуры Employee и Department.
// Employee должен содержать поля Name, ID и Department (вложенная структура Department).
// Напишите функцию, которая выводит информацию о сотруднике и его отделе.
// package main

// import "fmt"
// type Employee struct{
// 	Name string
// 	ID	string
// 	Department Department
// }
// type Department struct{
// 	Name string
// }
// func printEmployee(e Employee)  {
// 	fmt.Printf("Name: %s, ID: %s\n",e.Name,e.ID)
// 	fmt.Printf("Department: %s\n",e.Department.Name)
// }
// func main() {
// 	printEmployee(Employee{Name: "Halid",ID: "21",Department: Department{Name: "Finance"}})
// }

// -------------------------------------------------------------------------------------------------
// Дерево сотрудников
// Создайте структуру Employee с полем Manager,
// которое указывает на другого сотрудника (вложенная структура).
// Напишите функцию, которая рекурсивно выводит информацию о сотруднике и его менеджере.
// package main

// import "fmt"
// type Employee struct{
// 	Name string
// 	Position	string
// 	Manager *Employee
// }

// func printEmployee(e *Employee)  {
// 	if e==nil{
// 		return
// 	}
// 		fmt.Printf("Name: %s, Position: %s\n",e.Name,e.Position)
// 		if e.Manager!=nil{
// 			fmt.Printf("Manager: %s\n",e.Manager.Name)
// 			printEmployee(e.Manager)

//		}
//	}
//
//	func main() {
//		printEmployee(&Employee{Name: "Salim",Position: "1",Manager: &Employee{Name: "Hotam",Position: "2",Manager: &Employee{Name: "Halim",Position: "3"}}})
//	}
//
// -------------------------------------------------------------------------------------------------
// Временная структура для хранения информации о книге
// Создайте анонимную структуру для хранения информации о книге с полями Title и Author.
// Напишите функцию, которая принимает эту анонимную структуру и выводит информацию о книге.
// package main

// import "fmt"
// type book struct{
// 	Title string
// 	Author string
// }

// func printBook(book *book)  {
// 	fmt.Printf("Title: %s\nAuthor: %s\n",book.Title,book.Author)
// }
// func main() {
// 	bookStore:=new(book)
// 	bookStore.Title="Sad music"
// 	bookStore.Author="Mee Dun"

//		printBook(bookStore)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"
// func printBook(book struct{Title string;Author string}) {
// 	fmt.Printf("Title: %s\nAuthor: %s\n",book.Title,book.Author)
// }
// func main() {
// 	printBook(struct{Title string; Author string}{Title: "Sad music",Author: "Man Dun"})
// }
// -------------------------------------------------------------------------------------------------
// Временная структура для хранения информации о продукте
// Создайте анонимную структуру для хранения информации о продукте с полями Name и Price. Напишите функцию,
// которая принимает эту анонимную структуру и возвращает строку с описанием продукта.
// package main

// import (
//     "fmt"
//     "strconv"
// )

// // Функция для создания строки описания продукта
// func describeProduct(product struct {
//     Name  string
//     Price float64
// }) string {
//     return "Product Name: " + product.Name + ", Price: $" + strconv.FormatFloat(product.Price, 'f', 2, 64)
// }

// func main() {
//     // Создание анонимной структуры с информацией о продукте
//     product := struct {
//         Name  string
//         Price float64
//     }{
//         Name:  "Widget",
//         Price: 19.99,
//     }

//	    // Использование функции для получения строки описания продукта
//	    description := describeProduct(product)
//	    fmt.Println(description)
//	}
//
// -------------------------------------------------------------------------------------------------
// Временная структура для хранения информации о событии
// Создайте анонимную структуру для хранения информации о событии с полями EventName и Date.
// Напишите функцию, которая принимает эту анонимную структуру и возвращает строку с описанием события.
package main

import "fmt"

func evenPrint(event struct {
	EventName string
	Data      string
}) string {
	return fmt.Sprintf("EventName: %s\nData: %s\n",event.EventName,event.Data)
}

func main() {
	result:=evenPrint(struct {
		EventName string
		Data      string
	}{
		EventName: "Wedding",
		Data:      "21.12.22",
	})
	fmt.Print(result)
}
