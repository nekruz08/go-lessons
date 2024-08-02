// package main

// import "fmt"

// type MySlice []int // Sum

// // Метод для пользовательского типа MySlice
// func (s MySlice) Sum() int {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return sum
// }

//	func main() {
//		s := MySlice{1, 2, 3, 4, 5}
//		fmt.Println("Sum:", s.Sum()) // Output: Sum: 15
//	}
//
// -------------------------------------------------------------------------------------------------
// При встраивании структур, методы встроенных структур также могут быть
// использованы для удовлетворения интерфейсов.
// package main

// import "fmt"

// type Speaker interface {
// 	Speak()
// }
// type Animal struct {
// 	name string
// }

// func (a Animal) Speak() {
// 	fmt.Println(a.name, "says hello!")
// }

// type Dog struct {
// 	Animal
// 	breed string
// }

// func main() {
// 	var speaker Speaker
// 	dog := Dog{Animal: Animal{name: "Buddy"}, breed: "Golden Retriever"}
// 	speaker = dog
// 	speaker.Speak() // Output: Buddy says hello!

// }
// -------------------------------------------------------------------------------------------------
// Методы для пользовательских коллекций
// Вы можете создавать методы для пользовательских коллекций, таких как карты
// и срезы, для добавления специализированного поведения.

// package main

// import "fmt"

// type IntMap map[string]int

// // Sum
// // Метод для IntMap
//
//	func (m IntMap) Sum() int {
//		sum := 0
//		for _, v := range m {
//			sum += v
//		}
//		return sum
//	}
//
//	func main() {
//		m := IntMap{"a": 1, "b": 2, "c": 3}
//		fmt.Println("Sum:", m.Sum()) // Output: Sum: 6
//	}
//
// -------------------------------------------------------------------------------------------------
// Методы с вариативными параметрами
// Методы могут использовать вариативные параметры для обработки переменного
// числа аргументов.

// package main

// import "fmt"

// type Logger struct{}

// // Log
// // Метод с вариативными параметрами
// func (Logger) Log(messages ...string) {
// 	for _, message := range messages {
// 		fmt.Println("Log:", message)
// 	}
// }
// func main() {
// 	logger := Logger{}
// 	logger.Log("Message 1", "Message 2", "Message 3") // Output:	// Log: Message 1
// 	// Log: Message 2
// 	// Log: Message 3
// }
// -------------------------------------------------------------------------------------------------
// Методы и обработка ошибок
// Методы могут возвращать ошибки для обработки исключительных ситуаций.
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Divider struct{}

// // Divide
// // Метод для деления с обработкой ошибок
// func (Divider) Divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	div := Divider{}
// 	result, err := div.Divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err) // Output: Error: division by zero
// 	} else {
// 		fmt.Println("Result:", result)
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Методы и конструкторы
// В Go нет явной поддержки конструкторов, но вы можете использовать функции
// для инициализации структур.

// package main

// import "fmt"

// type Rectangle struct {
// 	width, height float64
// }

// // NewRectangle
// // Функция-конструктор для Rectangle
// func NewRectangle(width, height float64) *Rectangle {
// 	return &Rectangle{width: width, height: height}
// }
// func main() {
// 	rect := NewRectangle(10, 5)
// 	fmt.Printf("Rectangle: %+v\n", rect) // Output: Rectangle: &{width:10 height:5}
// }
// -------------------------------------------------------------------------------------------------
// Методы и замыкания (closures)
// Методы могут использовать замыкания для хранения состояния и создания
// более сложного поведения.
package main

import "fmt"

type Counter struct {
	value int
}

// Incrementer 
// Метод, возвращающий функцию-замыкание
func (c *Counter) Incrementer() func() int {
	return func() int {
		c.value++
		return c.value
	}
}

func main() {
	counter := Counter{}
	inc := counter.Incrementer()

	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3
}