// Задания для 14 – го урока
// Подсчет символов в файле
// Описание: Напишите программу, которая считывает файл и подсчитывает количество символов в нём.
// Методы или функции:
// func countCharacters(fileName string) (int, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	count, err := countCharacters("example.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Number of characters:", count)
// }

// func countCharacters(fileName string) (int, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanRunes)

// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

//-------------------------------------------------------------------------------------------------
// Подсчет строк в файле
// Описание: Напишите программу, которая считает количество строк в текстовом файле.
// Методы или функции:
// func countLines(fileName string) (int, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	count, err := countCharacters("example.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Number of characters:", count)
// }

// func countCharacters(fileName string) (int, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanLines)

// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }
//-------------------------------------------------------------------------------------------------
// Счетчик слов в строке
// Описание: Напишите функцию, которая считает количество слов в строке.
// Методы или функции:
// func countWords(s string) int
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	count, err := countCharacters("example.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Number of characters:", count)
// }

// func countCharacters(fileName string) (int, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanWords)

// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return 0, err
// 	}

//		return count, nil
//	}
//
// -------------------------------------------------------------------------------------------------
// Запись строки в файл
// Описание: Напишите программу, которая записывает заданную строку в файл.
// Методы или функции:
// func writeStringToFile(fileName, content string) error
// package main

// import (
// 	"fmt"
// 	"os"
// 	"unicode/utf8"
// )

// func main() {
// 	fileName := "example.txt"
// 	content := "Hello, Go!"
// 	err := writeStringToFile(fileName, content)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	count, err := countRunes(fileName)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Number of runes:", count)
// }

// // writeStringToFile writes a string to a file
// func writeStringToFile(fileName, content string) error {
// 	file, err := os.Create(fileName) // Create a new file or truncate an existing file
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(content) // Write the string to the file
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // countRunes counts the number of runes in a file
// func countRunes(fileName string) (int, error) {
// 	data, err := os.ReadFile(fileName) // Read the contents of the file
// 	if err != nil {
// 		return 0, err
// 	}

//		return utf8.RuneCount(data), nil // Count the number of runes
//	}
//
// -------------------------------------------------------------------------------------------------
// Чтение первой строки файла
// Описание: Напишите программу, которая читает первую строку из текстового файла.
// Методы или функции:
// func readFirstLine(fileName string) (string, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fileName := "example.txt"
// 	firstLine, err := readFirstLine(fileName)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("First line:", firstLine)
// }

// // readFirstLine читает первую строку из файла
// func readFirstLine(fileName string) (string, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	if scanner.Scan() {
// 		return scanner.Text(), nil
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return "", err
// 	}

//		return "", nil // Возвращаем пустую строку, если файл пуст
//	}
//
// -------------------------------------------------------------------------------------------------
// Копирование содержимого одного файла в другой
// Описание: Напишите программу, которая копирует содержимое одного файла в другой.
// Методы или функции:
// func copyFile(src, dst string) error
// package main

// import (
// 	"io"
// 	"os"
// )

// func main() {
// 	copyFile("example.txt","copy.txt")
// }
// func copyFile(src, dst string) error {
// 	// Открываем исходный файл для чтения
// 	srcFile, err := os.Open(src)
// 	if err != nil {
// 		return err
// 	}
// 	defer srcFile.Close()

// 	// Создаем или открываем целевой файл для записи
// 	dstFile, err := os.Create(dst)
// 	if err != nil {
// 		return err
// 	}
// 	defer dstFile.Close()

// 	// Копируем содержимое исходного файла в целевой файл
// 	_, err = io.Copy(dstFile, srcFile)
// 	if err != nil {
// 		return err
// 	}

//		return nil
//	}
//
// -------------------------------------------------------------------------------------------------
// Чтение строки с консоли и запись в файл
// Описание: Напишите программу, которая читает строку с консоли и записывает её в файл.
// Методы или функции:
// func readAndWriteToFile(fileName string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	err := readAndWriteToFile("copy.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }

// // readAndWriteToFile читает строку с консоли и записывает её в файл
// func readAndWriteToFile(fileName string) error {
// 	// Открытие файла для записи (создание нового или перезапись существующего)
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Создание нового буферизованного читателя для чтения с консоли
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your name: ")

// 	// Чтение строки с консоли
// 	input, err := reader.ReadString('n')
// 	if err != nil {
// 		return err
// 	}

// 	// Запись строки в файл
// 	_, err = file.WriteString(input) // Записываем строку в файл, включая символ новой строки
// 	if err != nil {
// 		return err
// 	}

//		return nil
//	}
//
// -------------------------------------------------------------------------------------------------
// Обратное чтение файла
// Описание: Напишите программу, которая читает файл с конца в начало и выводит его содержимое на экран.
// Методы или функции:
// func reverseReadFile(fileName string) (string, error)
// package main

// import (
// 	"fmt"
// 	"io/os"
// 	"os"
// )

// func reverseReadFile(fileName string) (string, error) {
// 	// Открываем файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Читаем файл в буфер
// 	buffer, err := os.ReadAll(file)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Декодируем буфер в строку
// 	content := string(buffer)

// 	// Создаем новый срез рун для хранения результата
// 	reversed := make([]rune, len(content))

// 	// Копируем рунные элементы в обратном порядке
// 	for i, r := range content {
// 		reversed[len(content)-1-i] = r
// 	}

// 	reversedContent := string(reversed)
// 	return reversedContent, nil
// }

//	func main() {
//		fileName := "example.txt" // Замените на имя вашего файла
//		content, err := reverseReadFile(fileName)
//		if err != nil {
//			fmt.Println("Ошибка при чтении файла:", err)
//			return
//		}
//		fmt.Println("Содержимое файла (в обратном порядке):")
//		fmt.Println(content)
//	}
//
// -------------------------------------------------------------------------------------------------
// Объединение содержимого двух файлов
// Описание: Напишите программу, которая объединяет содержимое двух файлов и записывает результат в новый файл.
// Методы или функции:
// func concatenateFiles(file1, file2, outputFile string) error
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func concatenateFiles(file1, file2, outputFile string) error {
// 	// Открываем первый файл
// 	file1Content, err := os.ReadFile(file1)
// 	if err != nil {
// 		return err
// 	}

// 	// Открываем второй файл
// 	file2Content, err := os.ReadFile(file2)
// 	if err != nil {
// 		return err
// 	}

// 	// Объединяем содержимое файлов
// 	combinedContent := append(file1Content, file2Content...)

// 	// Записываем результат в новый файл
// 	err = os.WriteFile(outputFile, combinedContent, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Содержимое файлов %s и %s успешно объединено в файл %sn", file1, file2, outputFile)
// 	return nil
// }

// func main() {
// 	file1 := "first.txt"
// 	file2 := "second.txt"
// 	outputFile := "third.txt"

// 	err := concatenateFiles(file1, file2, outputFile)
// 	if err != nil {
// 		fmt.Println("Ошибка при объединении файлов:", err)
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Проверка существования файла
// Описание: Напишите функцию, которая проверяет, существует ли файл с заданным именем.
// Методы или функции:
// func fileExists(fileName string) bool
// package main

// import (
//  "fmt"
//  "os"
// )

// func fileExists(fileName string) bool {
//  _, err := os.Stat(fileName)
//  if err == nil {
//   return true // Файл существует
//  } else if os.IsNotExist(err) {
//   return false // Файл не существует
//  }
//  // Обработка других ошибок (например, нет доступа к файлу)
//  fmt.Println("Ошибка при проверке файла:", err)
//  return false
// }

//	func main() {
//	 fileName := "example.txt" // Замените на имя вашего файла
//	 if fileExists(fileName) {
//	  fmt.Printf("Файл %s существует.n", fileName)
//	 } else {
//	  fmt.Printf("Файл %s не существует.n", fileName)
//	 }
//	}
//
// -------------------------------------------------------------------------------------------------
// Чтение и подсчет уникальных слов в файле
// Описание: Напишите программу, которая читает файл и подсчитывает количество уникальных слов.
// Методы или функции:
// func countUniqueWords(fileName string) (int, error)
// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func countUniqueWords(fileName string) (int, error) {
// 	// Читаем содержимое файла
// 	content, err := os.ReadFile(fileName)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Разделяем текст на слова
// 	words := strings.Fields(string(content))

// 	// Создаем множество для уникальных слов
// 	uniqueWords := make(map[string]bool)
// 	for _, word := range words {
// 		uniqueWords[word] = true
// 	}

// 	return len(uniqueWords), nil
// }

// func main() {
// 	fileName := "example.txt"
// 	uniqueWordCount, err := countUniqueWords(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при подсчете уникальных слов:", err)
// 		return
// 	}
// 	fmt.Printf("Количество уникальных слов в файле '%s': %dn", fileName, uniqueWordCount)
// }
// -------------------------------------------------------------------------------------------------
// Поиск и замена слова в файле
// Описание: Напишите программу, которая заменяет все вхождения определенного слова в файле на другое слово.
// Методы или функции:
// func replaceWordInFile(fileName, oldWord, newWord string) error
// package main

// import (
//  "fmt"
//  "os"
//  "strings"
// )

// func replaceWordInFile(fileName, oldWord, newWord string) error {
//  // Читаем содержимое файла
//  content, err := os.ReadFile(fileName)
//  if err != nil {
//   return err
//  }

//  // Заменяем все вхождения oldWord на newWord
//  updatedContent := strings.ReplaceAll(string(content), oldWord, newWord)

//  // Записываем обновленное содержимое обратно в файл
//  err = os.WriteFile(fileName, []byte(updatedContent), 0644)
//  if err != nil {
//   return err
//  }

//  fmt.Printf("Слово '%s' успешно заменено на '%s' в файле '%s'n", oldWord, newWord, fileName)
//  return nil
// }

// func main() {
//  fileName := "example.txt" // Замените на имя вашего файла
//  oldWord := "кухна"          // Замените на слово, которое нужно заменить
//  newWord := "нав"          // Замените на новое слово
//  err := replaceWordInFile(fileName, oldWord, newWord)
//  if err != nil {
//   fmt.Println("Ошибка при замене слова:", err)
//  }
// }
// -------------------------------------------------------------------------------------------------
// Подсчет частоты слов в файле
// Описание: Напишите программу, которая подсчитывает частоту каждого слова в файле.
// Методы или функции:
// func wordFrequency(fileName string) (map[string]int, error)
// package main

// import (
//  "fmt"
//  "os"
//  "strings"
// )

// func wordFrequency(fileName string) (map[string]int, error) {
//  // Читаем содержимое файла
//  content, err := os.ReadFile(fileName)
//  if err != nil {
//   return nil, err
//  }

//  // Разделяем текст на слова
//  words := strings.Fields(string(content))

//  // Создаем карту для подсчета частоты слов
//  wordCount := make(map[string]int)
//  for _, word := range words {
//   wordCount[word]++
//  }

//  return wordCount, nil
// }

// func main() {
//  fileName := "example.txt" // Замените на имя вашего файла
//  wordFreq, err := wordFrequency(fileName)
//  if err != nil {
//   fmt.Println("Ошибка при подсчете частоты слов:", err)
//   return
//  }

//	 fmt.Println("Частота слов в файле:")
//	 for word, freq := range wordFreq {
//	  fmt.Printf("%s: %d\n", word, freq)
//	 }
//	}
//
// -------------------------------------------------------------------------------------------------
// / Переворачивание строк в файле
// Описание: Напишите программу, которая переворачивает строки в файле и записывает их в новый файл.
// Методы или функции:
// func reverseLines(fileName, outputFile string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func reverseLines(fileName, outputFileName string) error {
// 	// Открываем входной файл для чтения
// 	inputFile, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer inputFile.Close()

// 	// Открываем выходной файл для записи
// 	outputFile, err := os.Create(outputFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()

// 	// Читаем строки из входного файла и сохраняем их в срез
// 	scanner := bufio.NewScanner(inputFile)
// 	lines := make([]string, 0)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	// Записываем строки в обратном порядке в выходной файл
// 	for i := len(lines) - 1; i >= 0; i-- {
// 		fmt.Fprintln(outputFile, lines[i])
// 	}

// 	return nil
// }

//	func main() {
//		fileName := "input.txt"        // Замените на имя вашего входного файла
//		outputFileName := "output.txt" // Замените на имя выходного файла
//		err := reverseLines(fileName, outputFileName)
//		if err != nil {
//			fmt.Println("Ошибка при переворачивании строк:", err)
//			return
//		}
//		fmt.Printf("Строки из файла '%s' успешно перевернуты и записаны в файл '%s'\n", fileName, outputFileName)
//	}
//
// -------------------------------------------------------------------------------------------------
// Удаление пустых строк из файла
// Описание: Напишите программу, которая удаляет все пустые строки из файла.
// Методы или функции:
// func removeEmptyLines(fileName, outputFile string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func removeEmptyLines(inputFileName, outputFileName string) error {
// 	// Открываем входной файл для чтения
// 	inputFile, err := os.Open(inputFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer inputFile.Close()

// 	// Открываем выходной файл для записи
// 	outputFile, err := os.Create(outputFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()

// 	// Читаем строки из входного файла и записываем их в выходной файл, если они не пустые
// 	scanner := bufio.NewScanner(inputFile)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line != "" {
// 			fmt.Fprintln(outputFile, line)
// 		}
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	inputFileName := "first.txt"        // Замените на имя вашего входного файла
// 	outputFileName := "second.txt"      // Замените на имя выходного файла
// 	err := removeEmptyLines(inputFileName, outputFileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при удалении пустых строк:", err)
// 		return
// 	}
// 	fmt.Printf("Пустые строки из файла '%s' успешно удалены и записаны в файл '%s'\n", inputFileName, outputFileName)
// }
// -------------------------------------------------------------------------------------------------
// Сравнение двух файлов на идентичность
// Описание: Напишите программу, которая сравнивает два файла и определяет, идентичны ли они.
// Методы или функции:
// func compareFiles(file1, file2 string) (bool, error)
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func compareFiles(file1, file2 string) (bool, error) {
// 	content1, err := os.ReadFile(file1)
// 	if err != nil {
// 		return false, err
// 	}

// 	content2, err := os.ReadFile(file2)
// 	if err != nil {
// 		return false, err
// 	}

// 	return string(content1) == string(content2), nil
// }

// func main() {
// 	file1 := "first.txt" // Замените на имя первого файла
// 	file2 := "second.txt" // Замените на имя второго файла

// 	equal, err := compareFiles(file1, file2)
// 	if err != nil {
// 		fmt.Println("Ошибка при сравнении файлов:", err)
// 		return
// 	}

//		if equal {
//			fmt.Printf("Файлы '%s' и '%s' идентичны.\n", file1, file2)
//		} else {
//			fmt.Printf("Файлы '%s' и '%s' не идентичны.\n", file1, file2)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Чтение файла с конца
// Описание: Напишите программу, которая читает файл с конца и выводит его содержимое на экран.
// Методы или функции:
// func readFileReverse(fileName string) error
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func readFileReverse(fileName string) (string, error) {
// 	// Открываем файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Читаем файл в буфер
// 	buffer, err := io.ReadAll(file)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Декодируем буфер в строку
// 	content := string(buffer)

// 	// Создаем новый срез рун для хранения результата
// 	reversed := make([]rune, len(content))

// 	// Копируем рунные элементы в обратном порядке
// 	for i, r := range content {
// 		reversed[len(content)-1-i] = r
// 	}

// 	reversedContent := string(reversed)
// 	return reversedContent, nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	content, err := readFileReverse(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}
// 	fmt.Println("Содержимое файла (в обратном порядке):")
// 	fmt.Println(content)
// }
// -------------------------------------------------------------------------------------------------
// Подсчет количества строк с определенным словом
// Описание: Напишите программу, которая считает количество строк в файле, содержащих определенное слово.
// Методы или функции:
// func countLinesWithWord(fileName, word string) (int, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func countLinesWithWord(fileName, word string) (int, error) {
// 	// Открываем файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer file.Close()

// 	// Создаем новый сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	lineCount := 0

// 	// Перебираем строки в файле
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if strings.Contains(line, word) {
// 			lineCount++
// 		}
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return 0, err
// 	}

// 	return lineCount, nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	word := "apple"           // Замените на нужное слово
// 	count, err := countLinesWithWord(fileName, word)
// 	if err != nil {
// 		fmt.Println("Ошибка при подсчете строк:", err)
// 		return
// 	}
// 	fmt.Printf("Количество строк с словом '%s' в файле '%s': %d\n", word, fileName, count)
// }
// -------------------------------------------------------------------------------------------------
// Генерация файла с повторяющимися строками
// Описание: Напишите программу, которая генерирует файл, содержащий заданную строку, повторенную указанное количество раз.
// Методы или функции:
// func generateRepeatedLinesFile(fileName, line string, count int) error
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func generateRepeatedLinesFile(fileName, line string, count int) error {
// 	// Создаем файл для записи
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Записываем строку в файл заданное количество раз
// 	for i := 0; i < count; i++ {
// 		_, err := fmt.Fprintln(file, line)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func main() {
// 	fileName := "output.txt" // Замените на имя файла, который вы хотите создать
// 	line := "Hello, world!"  // Замените на вашу строку
// 	count := 10              // Замените на количество повторений

// 	err := generateRepeatedLinesFile(fileName, line, count)
// 	if err != nil {
// 		fmt.Println("Ошибка при генерации файла:", err)
// 		return
// 	}
// 	fmt.Printf("Файл '%s' успешно создан с %d повторениями строки '%s'\n", fileName, count, line)
// }
// -------------------------------------------------------------------------------------------------
// Подсчет количества символов в каждой строке файла
// Описание: Напишите программу, которая подсчитывает количество символов в каждой строке файла и выводит их на экран.
// Методы или функции:
// func countCharactersPerLine(fileName string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func countCharactersPerLine(fileName string) error {
// 	// Открываем файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Printf("Строка: %s, Количество символов: %d\n", line, len(line))
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	err := countCharactersPerLine(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при подсчете символов:", err)
// 		return
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Поиск самого длинного слова в файле
// Описание: Напишите программу, которая находит самое длинное слово в файле и выводит его на экран.
// Методы или функции:
// func findLongestWord(fileName string) (string, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func findLongestWord(fileName string) (string, error) {
// 	// Открываем файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	var longestWord string

// 	for scanner.Scan() {
// 		// Разделяем строку на слова
// 		words := strings.Fields(scanner.Text())
// 		for _, word := range words {
// 			if len(word) > len(longestWord) {
// 				longestWord = word
// 			}
// 		}
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return "", err
// 	}

// 	return longestWord, nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	longestWord, err := findLongestWord(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}

// 	fmt.Println("Самое длинное слово в файле:", longestWord)
// }
// -------------------------------------------------------------------------------------------------
// Подсчет частоты встречаемости символов в файле
// Описание: Напишите программу, которая подсчитывает частоту встречаемости каждого символа в файле.
// Методы или функции:
// func charFrequency(fileName string) (map[rune]int, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func charFrequency(fileName string) (map[rune]int, error) {
// 	// Открываем файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	// Создаем карту для хранения частоты символов
// 	frequency := make(map[rune]int)

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		for _, char := range line {
// 			frequency[char]++
// 		}
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}

// 	return frequency, nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	frequency, err := charFrequency(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}

// 	fmt.Println("Частота встречаемости символов в файле:")
// 	for char, count := range frequency {
// 		fmt.Printf("%c: %d\n", char, count)
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Реверсирование слов в каждой строке файла
// Описание: Напишите программу, которая реверсирует слова в каждой строке файла и записывает результат в новый файл.
// Методы или функции:
// func reverseWordsInFile(fileName, outputFile string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func reverseWordsInFile(fileName, outputFile string) error {
// 	// Открываем входной файл для чтения
// 	inputFile, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer inputFile.Close()

// 	// Открываем выходной файл для записи
// 	outFile, err := os.Create(outputFile)
// 	if err != nil {
// 		return err
// 	}
// 	defer outFile.Close()

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(inputFile)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// Разделяем строку на слова
// 		words := strings.Fields(line)
// 		// Создаем срез для хранения перевернутых слов
// 		reversedWords := make([]string, len(words))
// 		// Переворачиваем порядок слов
// 		for i, word := range words {
// 			reversedWords[len(words)-1-i] = word
// 		}
// 		// Объединяем перевернутые слова в строку и записываем в выходной файл
// 		reversedLine := strings.Join(reversedWords, " ") + "\n"
// 		_, err := outFile.WriteString(reversedLine)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	fileName := "example.txt"     // Замените на имя вашего входного файла
// 	outputFileName := "output.txt" // Замените на имя выходного файла
// 	err := reverseWordsInFile(fileName, outputFileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при обработке файла:", err)
// 		return
// 	}

// 	fmt.Println("Реверсирование слов в каждой строке файла завершено.")
// }
// -------------------------------------------------------------------------------------------------
// Подсчет количества слов в каждой строке файла
// Описание: Напишите программу, которая подсчитывает количество слов в каждой строке файла и выводит результат на экран.
// Методы или функции:
// func countWordsPerLine(fileName string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func countWordsPerLine(fileName string) error {
// 	// Открываем файл для чтения
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// Разделяем строку на слова
// 		words := strings.Fields(line)
// 		// Подсчитываем количество слов
// 		wordCount := len(words)
// 		// Выводим количество слов в строке
// 		fmt.Printf("Количество слов в строке: %d\n", wordCount)
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	err := countWordsPerLine(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Найти и заменить слово в файле, с сохранением регистра
// Описание: Напишите программу, которая находит и заменяет все вхождения слова в файле, сохраняя регистр (с учетом заглавных и строчных букв).
// Методы или функции:
// func replaceWordWithCase(fileName, oldWord, newWord string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func replaceWordWithCase(fileName, oldWord, newWord string) error {
// 	// Открываем файл для чтения
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	var updatedLines []string

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// Заменяем старое слово на новое
// 		updatedLine := strings.ReplaceAll(line, oldWord, newWord)
// 		updatedLines = append(updatedLines, updatedLine)
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	// Создаем выходной файл для записи результатов
// 	outputFileName := "output.txt" // Замените на имя выходного файла
// 	outputFile, err := os.Create(outputFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()

// 	// Записываем обновленные строки в выходной файл
// 	for _, updatedLine := range updatedLines {
// 		_, err := outputFile.WriteString(updatedLine + "\n")
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	fmt.Printf("Замена слова \"%s\" на \"%s\" выполнена. Результат сохранен в файле %s.\n", oldWord, newWord, outputFileName)
// 	return nil
// }

// func main() {
// 	fileName := "input.txt" // Замените на имя вашего входного файла
// 	oldWord := "old_word"   // Замените на искомое слово
// 	newWord := "new_word"   // Замените на новое слово
// 	err := replaceWordWithCase(fileName, oldWord, newWord)
// 	if err != nil {
// 		fmt.Println("Ошибка при обработке файла:", err)
// 		return
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Поиск самого короткого слова в файле
// Описание: Напишите программу, которая находит самое короткое слово в файле и выводит его на экран.
// Методы или функции:
// func findShortestWord(fileName string) (string, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func findShortestWord(fileName string) (string, error) {
// 	// Открываем файл для чтения
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	// Создаем сканер для чтения файла
// 	scanner := bufio.NewScanner(file)
// 	var shortestWord string

// 	for scanner.Scan() {
// 		// Получаем все слова из текущей строки
// 		words := strings.Fields(scanner.Text())
// 		for _, word := range words {
// 			// Определяем самое короткое слово
// 			if shortestWord == "" || len(word) < len(shortestWord) {
// 				shortestWord = word
// 			}
// 		}
// 	}

// 	// Проверяем на ошибки, произошедшие во время сканирования
// 	if err := scanner.Err(); err != nil {
// 		return "", err
// 	}

// 	return shortestWord, nil
// }

// func main() {
// 	fileName := "your_file.txt" // Замените на имя вашего файла
// 	shortestWord, err := findShortestWord(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}

// 	fmt.Println("Самое короткое слово в файле:", shortestWord)
// }
// -------------------------------------------------------------------------------------------------
// Чтение и запись файла по частям
// Описание: Напишите программу, которая читает файл по частям (например, по 1024 байта) и записывает его содержимое в другой файл.
// Методы или функции:
// func copyFileInChunks(srcFileName, dstFileName string) error
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func copyFileInChunks(srcFileName, dstFileName string) error {
// 	// Открываем исходный файл для чтения
// 	srcFile, err := os.Open(srcFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer srcFile.Close()

// 	// Создаем целевой файл для записи
// 	dstFile, err := os.Create(dstFileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer dstFile.Close()

// 	bufferSize := 1024 // Размер буфера (в байтах)
// 	buffer := make([]byte, bufferSize)

// 	for {
// 		n, err := srcFile.Read(buffer)
// 		if err != nil {
// 			if err == io.EOF {
// 				break // Достигнут конец файла
// 			}
// 			return err
// 		}

// 		_, err = dstFile.Write(buffer[:n])
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	fmt.Printf("Файл %s успешно скопирован в %s.\n", srcFileName, dstFileName)
// 	return nil
// }

// func main() {
// 	srcFileName := "input.txt"  // Замените на имя вашего входного файла
// 	dstFileName := "output.txt" // Замените на имя выходного файла
// 	err := copyFileInChunks(srcFileName, dstFileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при копировании файла:", err)
// 		return
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Подсчет количества символов, слов и строк в файле
// Описание: Напишите программу, которая подсчитывает количество символов, слов и строк в файле и выводит результат на экран.
// Методы или функции:
// func countFileStats(fileName string) (int, int, int, error)
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// // countFileStats читает файл и возвращает количество символов, слов и строк.
// func countFileStats(fileName string) (int, int, int, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return 0, 0, 0, err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var charCount, wordCount, lineCount int

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		charCount += len(line)
// 		wordCount += len(strings.Fields(line))
// 		lineCount++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return 0, 0, 0, err
// 	}

// 	return charCount, wordCount, lineCount, nil
// }

// func main() {
// 	fileName := "example.txt" // Замените на имя вашего файла
// 	charCount, wordCount, lineCount, err := countFileStats(fileName)
// 	if err != nil {
// 		fmt.Println("Ошибка при чтении файла:", err)
// 		return
// 	}

// 	fmt.Printf("Символов: %d\nСлов: %d\nСтрок: %d\n", charCount, wordCount, lineCount)
// }
// -------------------------------------------------------------------------------------------------
// Поиск и удаление строки в файле
// Описание: Напишите программу, которая находит и удаляет строку с определенным содержимым из файла.
// Методы или функции:
// func deleteLineFromFile(fileName, lineToDelete string) error
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func deleteLineFromFile(fileName, lineToDelete string) error {
// 	// Открываем исходный файл
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	var updatedLines []string
// 	scanner := bufio.NewScanner(file)

// 	// Читаем строки и добавляем те, которые не равны строке для удаления
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line != lineToDelete {
// 			updatedLines = append(updatedLines, line)
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	// Создаем новый файл для записи обновленных строк
// 	outputFile, err := os.Create("output.txt") // Замените на имя выходного файла
// 	if err != nil {
// 		return err
// 	}
// 	defer outputFile.Close()

// 	// Записываем обновленные строки в новый файл
// 	for _, updatedLine := range updatedLines {
// 		_, err := outputFile.WriteString(updatedLine + "\n")
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	fmt.Printf("Строка \"%s\" успешно удалена из файла.\n", lineToDelete)
// 	return nil
// }

//	func main() {
//		fileName := "input.txt"      // Замените на имя вашего файла
//		lineToDelete := "delete me" // Замените на строку, которую нужно удалить
//		err := deleteLineFromFile(fileName, lineToDelete)
//		if err != nil {
//			fmt.Println("Ошибка при обработке файла:", err)
//			return
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Сортировка строк в файле в алфавитном порядке
// Описание: Напишите программу, которая сортирует строки в файле в алфавитном порядке и записывает результат в новый файл.
// Методы или функции:
// func sortFileLines(fileName, outputFile string) error
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sortFileLines(fileName, outputFile string) error {
	// Открываем входной файл для чтения
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("ошибка при открытии входного файла %s: %w", fileName, err)
	}
	defer file.Close()

	// Читаем строки из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Проверяем наличие ошибок при сканировании
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка при чтении входного файла: %w", err)
	}

	// Сортируем строки в алфавитном порядке
	sort.Strings(lines)

	// Открываем выходной файл для записи
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("ошибка при создании выходного файла %s: %w", outputFile, err)
	}
	defer outFile.Close()

	// Записываем отсортированные строки в выходной файл
	for _, line := range lines {
		_, err := outFile.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("ошибка при записи в выходной файл: %w", err)
		}
	}

	return nil
}

func main() {
	fileName := "input.txt"        // Замените на имя вашего входного файла
	outputFileName := "output.txt" // Замените на имя выходного файла

	err := sortFileLines(fileName, outputFileName)
	if err != nil {
		fmt.Println("Ошибка при сортировке строк:", err)
		return
	}

	fmt.Printf("Строки из файла '%s' успешно отсортированы и записаны в файл '%s'.\n", fileName, outputFileName)
}
