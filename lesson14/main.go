// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	bufer:=make([]byte,20)
// 	n,_:=file.Read(bufer)
// 	fmt.Println(string(bufer[:n]))

// 	fmt.Println("File opened successfully")

// }
//-------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// Открытие файла как источник потока данных
// 	file, err := os.Open("largefile.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Чтение файла и вывод его содержимого по частям
// 	buffer := make([]byte, 1024) // Чтение по 1024 байта
// 	for {
// 		n, err := file.Read(buffer)
// 		if err == io.EOF {
// 			break // Достигнут конец файла
// 		}
// 		if err != nil {
// 			fmt.Println("Error reading file:", err)
// 			return
// 		}
// 		// Вывод прочитанной части файла
// 		fmt.Print(string(buffer[:n]))
// 	}
// }
//-------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     var input string
//     fmt.Print("Enter your name: ")
//     fmt.Fscan(os.Stdin, &input)
//     fmt.Printf("Hello, %s!\n", input)
// }
//-------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     fmt.Fprintln(os.Stdout, "This is a message to stdout.")
// }
//-------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     fmt.Fprintln(os.Stderr, "This is an error message.")
// }
//-------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// Открытие файла для чтения
// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	// Копирование данных из файла в стандартный поток вывода
// 	_, err = io.Copy(os.Stdout, file)
// 	if err != nil {
// 		fmt.Println("Error copying data:", err)
// 		return
// 	}
// }
//-------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// Открытие файла-источника для чтения
// 	srcFile, err := os.Open("source.txt")
// 	if err != nil {
// 		fmt.Println("Error opening source file:", err)
// 		return
// 	}
// 	defer srcFile.Close()

// 	// Создание файла-назначения для записи
// 	dstFile, err := os.Create("destination.txt")
// 	if err != nil {
// 		fmt.Println("Error creating destination file:", err)
// 		return
// 	}
// 	defer dstFile.Close()

// 	// Копирование данных из файла-источника в файл-назначение
// 	bytesCopied, err := io.Copy(dstFile, srcFile)
// 	if err != nil {
// 		fmt.Println("Error copying data:", err)
// 		return
// 	}

//		fmt.Printf("Copied %d bytes from %s to %s.\n", bytesCopied, "source.txt", "destination.txt")
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Пример строки с Unicode символами
// 	content := "Hello, мир! 123"

// 	// Преобразуем строку в срез рун
// 	runes := []rune(content)

// 	// Разворачиваем срез рун
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}

// 	// Преобразуем срез рун обратно в строку
// 	reversedContent := string(runes)

//		// Выводим результаты
//		fmt.Println("Оригинальная строка:", content)
//		fmt.Println("Развернутая строка:", reversedContent)
//	}
//
// -------------------------------------------------------------------------------------------------
package main

import (
	"fmt"
)

// Функция для разворота среза рун
func reverseRunes(runes []rune) []rune {
	// Создаем новый срез рун для хранения результата
	reversed := make([]rune, len(runes))

	// Копируем рунные элементы в обратном порядке
	for i, r := range runes {
		reversed[len(runes)-1-i] = r
	}
	return reversed
}

func main() {
	// Пример строки с Unicode символами
	content := "Hello, мир! 123"

	// Преобразуем строку в срез рун
	runes := []rune(content)

	// Разворачиваем срез рун с использованием функции
	reversedRunes := reverseRunes(runes)

	// Преобразуем срез рун обратно в строку
	reversedContent := string(reversedRunes)

	// Выводим результаты
	fmt.Println("Оригинальная строка:", content)
	fmt.Println("Развернутая строка:", reversedContent)
}
