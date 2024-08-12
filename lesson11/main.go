// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// var str1 string="Hello, World!"
// 	// fmt.Println(str1)

// 	// str2:="Hello, World!"
// 	// fmt.Println(str2)

// 	// str3:=`	Hello, World!
// 	// Hello, World!`
// 	// fmt.Println(str3)

// 	// s1:="Hello, "
// 	// s2:="World"
// 	// s3:=s1+s2
// 	// fmt.Println(s3)

// 	// s := "Hello"
// 	// c := s[1]
// 	// fmt.Println(c)         //101 байт с кодом
// 	// fmt.Println(string(c)) // 'e' приведение байт кодоа в string

// 	// s:="Hello, World!"
// 	// sub:=s[7:12] // поучаем значение по этим индексам т.е. тсроку World  12 не вклюсительно
// 	// fmt.Println(sub)

// 	// s1:="abc"
// 	// s2:="def"
// 	// fmt.Println(s1==s2)
// 	// fmt.Println(s1<s2)

// 	// s:="Hello, World!"
// 	// fmt.Println(strings.Contains(s,"World"))
// 	// fmt.Println(strings.Index(s,"World"))

// 	// s := "Hello, 世界"
// 	// bytes := []byte(s)
// 	// fmt.Println(bytes)

// 	// rune := []rune(s)
// 	// fmt.Println(rune)

// 	s := "Hello, 世界"
// 	fmt.Println(len(s)) // 13

// 	fmt.Println(utf8.RuneCountInString(s)) // 9 (при условии что счет начинается с нуля)

//		for i, r := range s {
//			fmt.Printf("%d: %c\n", i, r)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// Пример строки с различными символами
// 	str := "Aé世􏰀"
// 	for len(str) > 0 {
// 		// Получаем первый байт
// 		r, size := utf8.DecodeRuneInString(str)
// 		fmt.Println(r)

// 		fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)
// 		// Обрезаем строку, удаляя первый символ

// 		str = str[size:]
// 	}
// }
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	str := "Hello, 世界"

// 	for len(str) > 0 {
// 		r, size := utf8.DecodeLastRuneInString(str)
// 		fmt.Printf("%X %v\n", r, size)

//			str = str[:len(str)-size]
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {

//		str:="ААА"
//		for _, v := range str {
//			fmt.Println(v)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {

// 	str := "abc"
// 	fmt.Printf("%T\n",str[0])

//		for _, v := range str {
//			fmt.Printf("%T\n",v)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// Пример строки с различными символами
// 	str := "Aé世􏰀"
// 	for len(str) > 0 {
// 		// Получаем первый байт
// 		r, size := utf8.DecodeRuneInString(str)
// 		fmt.Println(r)

// 		fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)
// 		// Обрезаем строку, удаляя первый символ

//			str = str[size:]
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// Пример байтовой последовательности для символов "Aé世􏰀"

// 	data := []byte{0x41, 0xC3, 0xA9, 0xE4, 0xB8, 0x96, 0xF0, 0x9F, 0x98, 0x80}
// 	for len(data) > 0 {
// 		// Декодируем первую руну
// 		r, size := utf8.DecodeRune(data)
// 		fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)
// 		// Обрезаем байтовый массив, удаляя первый символ
// 		data = data[size:]
// 	}
// }
// -------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
// )

// func main() {
//     s := "Hello, 世界"

//     // Преобразование строки в байты
//     bytes := []byte(s)
//     fmt.Println(bytes) // [72 101 108 108 111 44 32 228 184 150 231 149 140]

//	    // Преобразование строки в руны
//	    runes := []rune(s)
//	    fmt.Println(runes) // [72 101 108 108 111 44 32 19990 30028]
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	str := "A"
// 	fmt.Printf("%X\n", str[0]) // в шестнадцаричной системе исчисление 65=0041
// 	fmt.Printf("%d\n", str[0])  // в десятчной  системе исчисление 65
// 	fmt.Printf("%b\n", str[0])  // в двоичной  системе исчисление 65

// }
// -------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
//     "unicode/utf8"
// )

// func main() {
//     // Пример строки с различными символами
//     str := "Aé世😀"

//     for len(str) > 0 {
//         // Получаем первый байт
//         r, size := utf8.DecodeRuneInString(str)

//         fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)

//	        // Обрезаем строку, удаляя первый символ
//	        str = str[size:]
//	    }
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// Пример строки с различными символами
// 	str := "Aé世😀"

// 	for len(str) > 0 {
// 		// Получаем первый байт
// 		r, size := utf8.DecodeRuneInString(str)

// 		// fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size) 		// ШЕСТНАДЦАРИЧНЫЙ
// 		// fmt.Printf("Руна: %c, Demical: %d, Занимает %d байт\n", r, r, size)			// ДЕСЯТЕРИЧНЫЙ
// 		fmt.Printf("Руна: %c, Binary: %b, Занимает %d байт\n", r, r, size) // ДВОИЧНЫЙ

// 		// Обрезаем строку, удаляя первый символ
// 		str = str[size:]
// 	}
// }

// -------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
//     "strconv"
// )

//	func main() {
//	    hexNumber := "1F600"
//	    decimalValue, _ := strconv.ParseInt(hexNumber, 16, 64)
//	    binaryNumber := strconv.FormatInt(decimalValue, 2)
//	    fmt.Println(binaryNumber)
//	}
//
// -------------------------------------------------------------------------------------------------
// Этот код декодирует байты обратно в символы и показывает, сколько байтов занимает
// каждый символ.
// Кодировка символа по байтам
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	// Пример байтовой последовательности для символов "Aé世😀"
// 	data := []byte{0x41, 0xC3, 0xA9, 0xE4, 0xB8, 0x96, 0xF0, 0x9F, 0x98, 0x80}
// 	// fmt.Println(len(data))

// 	for len(data) > 0 {
// 		// Декодируем первую руну
// 		r, size := utf8.DecodeRune(data)

// 		fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)

//			// Обрезаем байтовый массив, удаляя первый символ
//			data = data[size:]
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World!"
//		contains := strings.Contains(s, "World")
//		fmt.Println(contains) // true
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, Hello, Hello!"
//		count := strings.Count(s, "Hello")
//		fmt.Println(count) // 3
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World!"
//		hasPrefix := strings.HasPrefix(s, "Hello") // Проверяет только перфик, суфикс не смотрить !
//		fmt.Println(hasPrefix) // true
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World!"
//		hasSuffix := strings.HasSuffix(s, "World!")
//		fmt.Println(hasSuffix) // true
//	}
//
// -------------------------------------------------------------------------------------------------
// Index
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World!"
//		index := strings.Index(s, "World")
//		fmt.Println(index) // 7
//	}
//
// -------------------------------------------------------------------------------------------------
// и LastIndex
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World!"
//		lastIndex := strings.LastIndex(s, "l") // Last index т.е проверяеть индекс последного l -a а не инедкс первых двух
//		fmt.Println(lastIndex) // 10
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

//		parts := []string{"Hello", "World"}
//		joined := strings.Join(parts, ", ")
//		fmt.Println(joined) // "Hello, World"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Go"
//		repeated := strings.Repeat(s, 3) //получаем новый стринг у утроенным Go, т.е. GoGoGo
//		fmt.Println(repeated) // "GoGoGo"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World! Hello, Go!"
//		// replaced := strings.Replace(s, "Hello", "Hi", 0) // Если нол то не работает
//		// replaced := strings.Replace(s, "Hello", "Hi", 1) // Если 1 то заменять будеть 1 раза
//		replaced := strings.Replace(s, "Hello", "Hi", 2) // Если 2 то заменять будеть 2 раза
//		fmt.Println(replaced)                            // "Hi, World! Hello, Go!"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

//	func main() {
//		s := "Hello, World! Hello, Go!"
//		replacedAll := strings.ReplaceAll(s, "Hello", "Hi")	// тут заменяется все
//		fmt.Println(replacedAll) // "Hi, World! Hi, Go!"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "a,b,c,d"
// 	split := strings.Split(s, ",") // отдаем оринтир по какому символу разделяй
// 	fmt.Println(split) // ["a", "b", "c", "d"]
// }
// -------------------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "a,b,c,d"
// 	splitN := strings.SplitN(s, ",", 2)  // использует запятую и делить на два (слай из двух строк)
// 	for _, v := range splitN {
// 		fmt.Println(v) // ["a", "b,c,d"]

//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "Hello, World!"
// 	lower := strings.ToLower(s)
// 	fmt.Println(lower) // "hello, world!"

// }
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "Hello, World!"
// 	upper := strings.ToUpper(s)
// 	fmt.Println(upper) // "HELLO, WORLD!"

// }
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "  Hello, World!  "
// 	// Функция strings.Trim удаляет пробелы в начале и в конце строки s.
// 	trimmed := strings.Trim(s, " ")
// 	fmt.Println(trimmed) // "Hello, World!"

//		// Функция strings.TrimSpace удаляет пробелы в начале и в конце строки s.
//		trimSpace := strings.TrimSpace(s)
//		fmt.Println(trimSpace) // "Hello, World!"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "Hello, World!"
// 	prefix := "Hello"

//		// Функция strings.TrimPrefix удаляет префикс prefix из строки s.
//		trimPrefix := strings.TrimPrefix(s, prefix)
//		fmt.Println(trimPrefix) // "  Hello, World!  "
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "Hello, World!"
// 	suffix := "World!"

//		// Функция strings.TrimSuffix удаляет суффикс suffix из строки s.
//		trimSuffix := strings.TrimSuffix(s, suffix)
//		fmt.Println(trimSuffix) // "  Hello, World!  "
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// Преобразование строки в целое число
// 	i, err := strconv.Atoi("123")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("%T,%d\n", i, i) // "123"
// 	}
// }

// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// Преобразование целого числа в строку
// 	s := strconv.Itoa(123)
// 	fmt.Printf("%T,%s",s,s) // "123"

// }
// -------------------------------------------------------------------------------------------------
// package main
// import (
//     "fmt"
//     "strconv"
// )

// func main() {
//     // Преобразование строки в целое число с указанием базы
//     i, err := strconv.ParseInt("1A", 16, 64)
//     if err != nil {
//         fmt.Println(err)
//     } else {
//         fmt.Println(i) // 26
//     }

//     // Преобразование целого числа в строку с указанием базы
//     s := strconv.FormatInt(26, 16)
//     fmt.Println(s) // "1a"
// }
// -------------------------------------------------------------------------------------------------
// package main
// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// Добавление кавычек к строке
// 	s := strconv.Quote("Hello, World!")
// 	fmt.Println(s) // "\"Hello, World!\""

//		// Удаление кавычек из строки
//		unquoted, err := strconv.Unquote(s)
//		if err != nil {
//			fmt.Println(err)
//		} else {
//			fmt.Println(unquoted) // "Hello, World!"
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

//	func main() {
//		b := []byte("Number: ")
//		b = strconv.AppendInt(b, 123, 10)
//		fmt.Println(string(b)) // "Number: 123"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode"
// )

// func main() {
// 	r := 'a' //it works just with one character
// 	fmt.Println(string(unicode.ToUpper(r))) // "A"

//		r = 'A'
//		fmt.Println(string(unicode.ToLower(r))) // "a"
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	b := []byte("世界")
// 	r, size := utf8.DecodeRune(b)
// 	fmt.Printf("Rune: %c, Size: %d\n", r, size) // Rune: H, Size: 1

// 	s := "世界"
// 	r, size = utf8.DecodeRuneInString(s)
// 	fmt.Printf("Rune: %c, Size: %d\n", r, size) // Rune: H, Size: 1
// }
// -------------------------------------------------------------------------------------------------
// package main
// import (
//     "fmt"
//     "unicode/utf8"
// )

//	func main() {
//	    r := '世'
//	    buf := make([]byte, 3)
//	    size := utf8.EncodeRune(buf, r)
//	    fmt.Printf("Encoded: % x, Size: %d\n", buf, size) // Encoded: e4 b8 96, Size: 3
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	r := 'A'
// 	size := utf8.RuneLen(r)
// 	fmt.Printf("Rune: %c, Size: %d\n", r, size) // Rune: 世, Size: 3
// }

// -------------------------------------------------------------------------------------------------
// package main
// import (
//     "fmt"
//     "unicode/utf8"
// )

// func main() {
//     b := []byte("Hello, 世界")
//     count := utf8.RuneCount(b)
//     fmt.Printf("Byte slice: %d runes\n", count) // Byte slice: 9 runes

//	    s := "Hello, 世界"
//	    count = utf8.RuneCountInString(s)
//	    fmt.Printf("String: %d runes\n", count) // String: 9 runes
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	b := []byte("Hello, 世界")
// 	count := utf8.RuneCount(b)
// 	fmt.Printf("Byte slice: %d runes\n", count) // Byte slice: 9 runes

// 	s := "Hello, 世界"
// 	count = utf8.RuneCountInString(s)
// 	fmt.Printf("String: %d runes\n", count) // String: 9 runes
// }

// -------------------------------------------------------------------------------------------------
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Hello, World!"

	// Получаем указатель на данные строки и её длину
	header := (*[2]uintptr)(unsafe.Pointer(&s))
	dataPtr := header[0]
	length := int(header[1])

	fmt.Printf("String: %s\n", s)
	fmt.Printf("Data pointer: %x\n", dataPtr)
	fmt.Printf("Length: %d\n", length)
}

// -------------------------------------------------------------------------------------------------
