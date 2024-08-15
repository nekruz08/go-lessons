// 11. Форматирование JSON-вывода
// Описание: Напишите программу, которая сериализует структуру в JSON с отступами для улучшения читаемости.
// Метод: Используйте json.MarshalIndent.
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// 12. Обновление значения в JSON
// Описание: Напишите программу, которая десериализует JSON в структуру, изменяет одно из значений, а затем сериализует структуру обратно в JSON. Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// 13. Парсинг JSON с разными типами данных
// Описание: Напишите программу, которая десериализует JSON-строку, содержащую данные разных типов (строки, числа, массивы, объекты), в map[string]interface{}.
// Пример данных:

// {
// "name": "John",
// "age": 30,
// "scores": [100, 98, 95],
// "address": {"city": "New York", "street": "5th Avenue"}
// }
// 14. Массив структур в JSON
// Описание: Напишите программу, которая сериализует массив структур Person в JSON и выводит результат на экран.
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// 15. Добавление нового объекта в JSON-массив
// Описание: Напишите программу, которая читает JSON-массив из файла, добавляет новый объект и записывает обновленный массив обратно в файл. Файл: people.json
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`

// }
// 16. Удаление объекта из JSON-массива
// Описание: Напишите программу, которая читает JSON-массив из файла, удаляет объект с определенным значением поля и записывает обновленный массив обратно в файл.
// Файл: people.json
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// 17. Поиск объекта в JSON-массиве
// Описание: Напишите программу, которая читает JSON-массив из файла и находит объект с определенным значением поля.
// Файл: people.json
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// 18. Подсчет объектов в JSON-массиве
// Описание: Напишите программу, которая читает JSON-массив из файла и

// подсчитывает количество объектов в массиве. Файл: people.json
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// 19. Десериализация JSON с произвольными полями
// Описание: Напишите программу, которая десериализует JSON-строку с произвольным набором полей в map[string]interface{}.
// Пример данных:
// {
// "name": "John", "age": 30, "country": "USA", "job": "Engineer"
// }
// 20. Сериализация и десериализация времени в JSON
// Описание: Напишите программу, которая сериализует и десериализует структуру с полем времени (time.Time) в JSON.
// Структура:
// type Event struct {
// Name string `json:"name"`

// Timestamp time.Time `json:"timestamp"` }
// -------------------------------------------------------------------------------------------------
// 20 JSON Tasks in Go
// 1. Сериализация структуры в JSON
// Описание: Напишите программу, которая сериализует структуру Person в формат JSON и выводит результат на экран.
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name    string  `json:"name"`
// 	Age     int     `json:"age"`
// 	Email   string  `json:"email"`
// }

// func main() {
// 	person := Person{
// 		Name:  "John",
// 		Age:   30,
// 		Email: "john.doe@example.com",
// 	}

// 	// Сериализация объекта в JSON
// 	jsonData, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("Error serializing to JSON:", err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }
// -------------------------------------------------------------------------------------------------
// 2. Десериализация JSON в структуру
// Описание: Напишите программу, которая десериализует JSON-строку в структуру Person и выводит результат на экран.
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"` Email string `json:"email"`
// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name    string  `json:"name"`
// 	Age     int     `json:"age"`
// 	Email   string  `json:"email"`
// }

// func main() {
// 	jsonData := `{
// 		"name": "John",
// 		"age": 30,
// 		"email": "john.doe@example.com"
// 	}`

// 	var person Person

// 	// Десериализация JSON-строки в объект
// 	err := json.Unmarshal([]byte(jsonData), &person)
// 	if err != nil {
// 		fmt.Println("Error deserializing JSON:", err)
// 		return
// 	}

// 	fmt.Printf("Name: %s, Age: %d, Email: %s\n",
// 		person.Name, person.Age, person.Email)
// }
// -------------------------------------------------------------------------------------------------
// 3. Сериализация карты в JSON
// Описание: Напишите программу, которая сериализует карту map[string]int в JSON и выводит результат на экран.
// Пример данных:
// data := map[string]int{ "apples": 5, "oranges": 10,
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	data := map[string]int{"apples": 5, "oranges": 10}

// 	// Сериализация объекта в JSON
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Error serializing to JSON:", err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }
// -------------------------------------------------------------------------------------------------
// 4. Десериализация JSON в карту
// Описание: Напишите программу, которая десериализует JSON-строку в карту map[string]int и выводит результат на экран.
// Пример данных:
// {
// "apples": 5,
// "oranges": 10 }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name    string  `json:"name"`
// 	Age     int     `json:"age"`
// 	Email   string  `json:"email"`
// 	Address Address `json:"address"`
// }

// type Address struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }

// func main() {
// 	jsonData := `{
// 		"apples": 5,
// 		"oranges": 10
// 	}`

// 	var data map[string]int

// 	// Десериализация JSON-строки в объект
// 	err := json.Unmarshal([]byte(jsonData), &data)
// 	if err != nil {
// 		fmt.Println("Error deserializing JSON:", err)
// 		return
// 	}

// 	for k, v := range data {
// 		fmt.Printf("%s: %v\n", k, v)
// 	}
// }
// -------------------------------------------------------------------------------------------------
// 5. Пропуск полей при сериализации
// Описание: Напишите программу, которая сериализует структуру, пропуская поля с пустыми значениями.
// Структура:
// type Product struct {
// Name string `json:"name"`
// Price float64 `json:"price,omitempty"` InStock bool `json:"in_stock,omitempty"`
// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Product struct {
// 	Name    string  `json:"name"`
// 	Price   float64 `json:"price,omitempty"`
// 	InStock bool    `json:"in_stock,omitempty"`
// }

// func main() {
// 	p1 := Product{
// 		Name:    "Laptop",
// 		Price:   999.99,
// 		InStock: true,
// 	}

// 	p2 := Product{
// 		Name: "Mouse",
// 	}

// 	data1, err := json.Marshal(p1)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON for p1:", err)
// 		return
// 	}

// 	data2, err := json.Marshal(p2)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON for p2:", err)
// 		return
// 	}

// 	fmt.Println("Product 1:", string(data1))
// 	fmt.Println("Product 2:", string(data2))
// }
// -------------------------------------------------------------------------------------------------
// 6. Работа с вложенными структурами
// Описание: Напишите программу, которая сериализует и десериализует структуру с вложенной структурой.
// Структура:

// type Address struct {
// Street string `json:"street"` City string `json:"city"`
// }
// type Person struct {
// Name string `json:"name"` Age int `json:"age"`
// Email string `json:"email"` Address Address `json:"address"`
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Address struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }
// type Person struct {
// 	Name    string  `json:"name"`
// 	Age     int     `json:"age"`
// 	Email   string  `json:"email"`
// 	Address Address `json:"address"`
// }

// func main() {
// 	address := Address{
// 		Street: "Mayakovskaya",
// 		City:   "Dushanbe",
// 	}

// 	person := Person{
// 		Name:    "Odam",
// 		Age:     21,
// 		Email:   "email@gmail.tj",
// 		Address: address,
// 	}

// 	serial, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON for p1:", err)
// 		return
// 	}
// 	fmt.Println("Сериализация:", string(serial))

// 	err = json.Unmarshal(serial, &person)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON for p2:", err)
// 		return
// 	}
// 	fmt.Printf("Десериализация: %+v", person)
// }
// -------------------------------------------------------------------------------------------------
// 7. Использование JSON с интерфейсами
// Описание: Напишите программу, которая сериализует и десериализует карту map[string]interface{}
// и выводит результат на экран.
// Пример данных:
// data := map[string]interface{}{ "name": "John",
// "age": 30,
// "email": "john@example.com",
// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	data := map[string]interface{}{
// 		"name":  "John",
// 		"age":   30,
// 		"email": "john@example.com",
// 	}

// 	serial, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON for p1:", err)
// 		return
// 	}
// 	fmt.Println("Сериализация:", string(serial))

// 	err = json.Unmarshal(serial, &data)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON for p2:", err)
// 		return
// 	}
// 	fmt.Printf("Десериализация: %+v", data)
// }
// -------------------------------------------------------------------------------------------------
// 8. Чтение JSON из файла
// Описание: Напишите программу, которая читает JSON из файла и десериализует его в структуру.
// Файл: person.json

// Структура:
// type Person struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// 	Email string `json:"email"`
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Person struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"email"`
// }

// func main() {
// 	// Открытие файла для чтения
// 	file, err := os.Open("person.json")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	var person Person

// 	// Создание нового JSON-декодера и чтение данных из файла
// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(&person)
// 	if err != nil {
// 		fmt.Println("Error decoding JSON from file:", err)
// 		return
// 	}

//		fmt.Printf("Name: %s, Age: %d, Email: %s\n", person.Name, person.Age, person.Email)
//	}
//
// -------------------------------------------------------------------------------------------------
// 9. Запись JSON в файл
// Описание: Напишите программу, которая сериализует структуру в JSON и записывает результат в файл.
// Файл: output.json
// Структура:
// type Product struct {
// Name string `json:"name"`
// Price float64 `json:"price"`
// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Product struct {
// 	Name  string  `json:"name"`
// 	Price float64 `json:"price"`
// }

// func main() {
// 	product := Product{
// 		Name:  "Apple",
// 		Price:   30,
// 	}

// 	// Открытие файла для записи
// 	file, err := os.Create("output.json")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Создание нового JSON-энкодера и запись данных в файл
// 	encoder := json.NewEncoder(file)
// 	err = encoder.Encode(product)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON to file:", err)
// 	}
// }
// -------------------------------------------------------------------------------------------------
// 10. Пропуск неизвестных полей при десериализации
// Описание: Напишите программу, которая десериализует JSON-строку в структуру, игнорируя неизвестные поля.
// Структура:
// type Person struct {
// Name string `json:"name"` Age int `json:"age"`
// }
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{
		"name": "Alice",
		"age": 30,
		"unknown_field": "This will be ignored",
		"abcd": "dfldas"
	}`

	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("Person: %+v\n", person)
}
