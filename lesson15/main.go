//-------------------------------------------------------------------------------------------------
// Сериализация: Преобразование объекта в JSON
// package main
// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	Email   string `json:"email"`
// 	Address Address `json:"address"`
// }

// type Address struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }

// func main() {
// 	person := Person{
// 		Name:  "John",
// 		Age:   30,
// 		Email: "john.doe@example.com",
// 		Address: Address{
// 			Street: "123 Main St",
// 			City:   "Springfield",
// 		},
// 	}

// 	// Сериализация объекта в JSON
// 	jsonData, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("Error serializing to JSON:", err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }
//-------------------------------------------------------------------------------------------------
// Десериализация: Преобразование JSON в объект
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	Email   string `json:"email"`
// 	Address Address `json:"address"`
// }

// type Address struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }

// func main() {
// 	jsonData := `{
// 		"name": "John",
// 		"age": 30,
// 		"email": "john.doe@example.com",
// 		"address": {
// 			"street": "123 Main St",
// 			"city": "Springfield"
// 		}
// 	}`

// 	var person Person

// 	// Десериализация JSON-строки в объект
// 	err := json.Unmarshal([]byte(jsonData), &person)
// 	if err != nil {
// 		fmt.Println("Error deserializing JSON:", err)
// 		return
// 	}

// 	fmt.Printf("Name: %s, Age: %d, Email: %s, Address: %s, %s\n",
// 		person.Name, person.Age, person.Email, person.Address.Street, person.Address.City)
// }
//-------------------------------------------------------------------------------------------------
// Запись объекта в JSON-файл
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	Email   string `json:"email"`
// 	Address Address `json:"address"`
// }

// type Address struct {
// 	Street string `json:"street"`
// 	City   string `json:"city"`
// }

// func main() {
// 	person := Person{
// 		Name:  "John",
// 		Age:   30,
// 		Email: "john.doe@example.com",
// 		Address: Address{
// 			Street: "123 Main St",
// 			City:   "Springfield",
// 		},
// 	}

// 	// Открытие файла для записи
// 	file, err := os.Create("person.json")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

//		// Создание нового JSON-энкодера и запись данных в файл
//		encoder := json.NewEncoder(file)
//		err = encoder.Encode(person)
//		if err != nil {
//			fmt.Println("Error encoding JSON to file:", err)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Чтение JSON из файла
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

func main() {
	// Открытие файла для чтения
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var person Person

	// Создание нового JSON-декодера и чтение данных из файла
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding JSON from file:", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d, Email: %s, Address: %s, %s\n",
		person.Name, person.Age, person.Email, person.Address.Street, person.Address.City)
}
