// Копирование map
// Описание: Напишите функцию, которая создает копию map.
// Преобразование map в срез пар ключ-значение
// Описание: Напишите функцию, которая преобразует map в срез пар ключ-значение.
// Обновление значений map на основе другой функции
// Описание: Напишите функцию, которая обновляет значения в map, применяя функцию к каждому значению.
// Кэширование с использованием map и sync.RWMutex
// Описание: Реализуйте простой кэш с использованием map и sync.RWMutex.
// Сложный фильтр по нескольким условиям
// Описание: Напишите функцию, которая фильтрует map по нескольким условиям.
// Слияние и агрегация значений из нескольких map
// Описание: Реализуйте функцию, которая сливает несколько map в один, суммируя значения для одинаковых ключей.
// Итерирование по ключам и значениям с использованием функции обратного вызова
// Описание: Напишите функцию, которая итерирует по map, вызывая функцию обратного вызова для каждого ключа и значения.
// Создание и использование карты с указателями на структуры
// Описание: Создайте map с указателями на структуры в качестве значений.
// Сопоставление map и массивов
// Описание: Напишите функцию, которая преобразует map в массив структур и наоборот.
// Реализация метода для структуры с map
// Описание: Реализуйте метод для структуры, который работает с map в этой структуре.
// Использование map для реализации простого хранилища данных
// Описание: Реализуйте простое хранилище данных с map для управления записями.
// Реализация кэширования с лимитом на количество элементов
// Описание: Реализуйте кэш с использованием map и ограничением на количество элементов.
// Реализация map для хранения срезов строк
// Описание: Реализуйте map, где ключи – это строки, а значения – срезы строк. Добавьте функции для добавления, удаления и получения срезов.

// Задания для 12 – го урока

// Создание и вывод map
// Описание: Создайте map для хранения строк и их длин,
// добавьте несколько элементов и выведите содержимое.
// package main

// import "fmt"

//	func main() {
//		m:=make(map[string]int)
//		m["a"]=1
//		m["b"]=2
//		m["c"]=3
//		for key, value := range m {
//			fmt.Printf("map[%s]=%d\n",key,value)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Проверка наличия ключа
// Описание: Создайте map с несколькими элементами и напишите функцию,
// которая проверяет наличие заданного ключа.
// package main

// import "fmt"

// func main() {
// 	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
// 	fmt.Println(check(m, "e"))
// }

//	func check(m map[string]int, a string) bool {
//		_, b := m[a]
//		return b
//	}
//
// -------------------------------------------------------------------------------------------------
// Обновление значения по ключу
// Описание: Напишите функцию для обновления значения в map по заданному ключу.
// package main

// import "fmt"

// func main() {
// 	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
// 	update(m, "e", 5)
// 	fmt.Println(m["e"])

// }

//	func update(mapp map[string]int, str string, num int) {
//		mapp[str] = num
//	}
//
// -------------------------------------------------------------------------------------------------
// Удаление элемента из map
// Описание: Создайте функцию, которая удаляет элемент из map по заданному ключу.
// package main

// import "fmt"

// func main() {
// 	m:=map[string]int{"a":1,"b":2}
// 	fmt.Println(m)
// 	deleteFunc(m,"b")
// 	fmt.Println(m)

// }
//
//	func deleteFunc(m map[string]int,key string)  {
//		delete(m,key)
//	}
//
// -------------------------------------------------------------------------------------------------
// Подсчет частоты слов
// Описание: Напишите функцию, которая подсчитывает частоту слов в строке и возвращает map с результатами.
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	s := "nekruz nekruz nekruz nekruz nekruz shodmon shodmon shodmon"
// 	// fmt.Println(strings.Count(s,"nekruz"))
// 	result := countWord(s)
// 	fmt.Println(result)

// }
// func countWord(s string) map[string]int {
// 	slice := strings.Fields(s)

//		mapp := make(map[string]int)
//		for _, v := range slice {
//			mapp[v]++
//		}
//		return mapp
//	}
//
// -------------------------------------------------------------------------------------------------
// Сортировка ключей в map
// Описание: Напишите функцию, которая сортирует ключи в map и возвращает отсортированные ключи.
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	m:=map[int]struct{}{1:{},2:{},3:{}}
// 	a(m)
// }

// func a(mapp map[int]struct{}) {
// 	sorted:=[]int{}
// 	for key,_:= range mapp {
// 		sorted = append(sorted, key)
// 	}

// 	less:= func(i, j int) bool{
// 		return i<j
// 	}

// 	sort.Slice(sorted,less)
// 	fmt.Println(sorted)
// }
// ----------------------------------------------
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	m:=map[string]struct{}{"c":{},"b":{},"a":{}}
// 	fmt.Println(a(m))
// }

//	func a(mapp map[string]struct{}) []string {
//		sorted:=make([]string,0,len(mapp))
//		for key:= range mapp {
//			sorted = append(sorted, key)
//		}
//		sort.Strings(sorted)
//		return sorted
//	}
//
// -------------------------------------------------------------------------------------------------
// Проверка пустого map
// Описание: Напишите функцию, которая проверяет, пустой ли map.
// package main

// import "fmt"

// func main() {
// 	var mapp map[string]int
// 	fmt.Println(checkNill(mapp))

// }
//
//	func checkNill(mapp map[string]int) bool {
//		return mapp == nil
//	}
//
// -------------------------------------------------------------------------------------------------
// Инвертирование ключей и значений
// Описание: Напишите функцию, которая инвертирует ключи и значения в map.
// package main

// import "fmt"

//	func main() {
//		mapp := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//		fmt.Println(invertMap(mapp))
//	}
//
//	func invertMap(m map[string]int) map[int]string {
//		inverted := make(map[int]string)
//		for key, value := range m {
//			inverted[value] = key
//		}
//		return inverted
//	}
//
// -------------------------------------------------------------------------------------------------
// Проверка дубликатов в map
// Описание: Напишите функцию, которая проверяет, есть ли дубликаты значений в map.
// package main

// import "fmt"

// func main() {
// 	mapp := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
// 	fmt.Println(invertMap(mapp))
// }

// func invertMap(m map[string]int) bool {
// 	inverted := make(map[int]bool)
// 	for _, value := range m {
// 		if inverted[value] {
// 			return true
// 		}
// 		inverted[value] = true

//		}
//		return false
//	}
//
// -------------------------------------------------------------------------------------------------
// Получение всех значений
// Описание: Напишите функцию, которая возвращает все значения из map в срезе.
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	mapp := map[string]int{"a": 1, "b": 2, "c": 3}

// 	fmt.Println(mapToSlice(mapp))

// }
// func mapToSlice(mapp map[string]int) []int {
// 	slice := make([]int, 0,len(mapp))
// 	for _, v := range mapp {
// 		slice = append(slice, v)
// 	}
// 	sort.Ints(slice)
// 	return slice

// }
// -------------------------------------------------------------------------------------------------
// Подсчет уникальных значений в срезе строк
// Описание: Напишите функцию, которая подсчитывает уникальные значения в срезе строк с использованием map.
// package main

// import "fmt"

// func main() {
// 	slice:=[]string{"a","b","c","d","b"}
// 	uniqeCount(slice)
// }
// func uniqeCount(slice []string)  {
// 	mapp:=make(map[string]int)
// 	for _, v := range slice {
// 		mapp[v]++
// 	}
// 	for key,value:= range mapp {
// 		fmt.Printf("%s = %v\n",key,value)
// 	}
// }
// ----------------------------------------------
// package main

// import "fmt"

//	func main() {
//		slice:=[]string{"a","b","c","d","b","c"}
//		uniqeCount(slice)
//	}
//
//	func uniqeCount(slice []string)  {
//		uniqValues:=make(map[string]struct{})
//		for _, v := range slice {
//			uniqValues[v]=struct{}{}
//		}
//		fmt.Println(len(uniqValues))
//	}
//
// -------------------------------------------------------------------------------------------------
// Поиск значений по маске
// Описание: Напишите функцию, которая возвращает все ключи, чьи значения удовлетворяют заданному условию.
// package main

// import "fmt"

// func main() {
// 	mapp:=map[string]int{"a":1,"b":2,"c":3,"d":4,"e":5}
// 	fmt.Println(sort(mapp))

// }
//
//	func sort(mapp map[string]int)[]string  {
//		newSlice:=make([]string,0)
//		for k, v := range mapp {
//			if v>=2{
//				newSlice = append(newSlice, k)
//			}
//		}
//		return newSlice
//	}
//
// -------------------------------------------------------------------------------------------------
// Объединение двух map
// Описание: Напишите функцию для объединения двух map.
// Если ключи совпадают, значения из второго map должны заменять значения из первого.
// package main

// import "fmt"

// func main() {
// 	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
// 	m2 := map[string]int{"f": 6, "g": 7, "c": 8, "i": 9, "j": 10}

// 	result:=combine2maps(m1, m2)
// 	for k, v := range result {
// 		fmt.Printf("%s=%v\n",k,v)
// 	}

// }
//
//	func combine2maps(map1,map2 map[string]int) map[string]int {
//		merged:=make(map[string]int)
//		for k, v := range map1 {
//			merged[k]=v
//		}
//		for k, v := range map2 {
//			merged[k]=v
//		}
//		return merged
//	}
//
// -------------------------------------------------------------------------------------------------
// Поиск по значению и получение ключа
// Описание: Напишите функцию, которая возвращает ключ по значению в map.
// Если значение не найдено, верните пустую строку.
// package main

// import "fmt"

// func main() {
// 	m:=make(map[string]int)
// 	m["a"]=1
// 	m["b"]=2
// 	m["c"]=3
// 	fmt.Println(a(3,m))

// }
// func a(v int,m map[string]int) (key string) {
// 	for key, value := range m {
// 		if value==v{
// 			return key
// 		}
// 	}
// 	return ""
// }
// -------------------------------------------------------------------------------------------------
// Пересечение двух map
// Описание: Напишите функцию для нахождения пересечения двух map на основе ключей и значений.

// package main

// func main() {
// 	first := map[string]int{"a": 1, "b": 1, "c": 8, "d": 1}
// 	second := map[string]int{"a": 1, "b": 1, "c": 8, "d": 1}
// 	intersectMaps(first, second)
// }

//	func intersectMaps(m1, m2 map[string]int) map[string]int {
//		intersected := make(map[string]int)
//		for key, value1 := range m1 {
//			if value2, exists := m2[key]; exists && value1 == value2 {
//				intersected[key] = value1
//			}
//		}
//		return intersected
//	}
//
// -------------------------------------------------------------------------------------------------
// Создание среза ключей из map
// Описание: Напишите функцию, которая возвращает срез всех ключей из map.
// package main

// import "fmt"

// func main() {
// 	m := map[string]int{"a": 1, "b": 1, "c": 8, "d": 1}
// 	fmt.Println(a(m))
// }
// func a(m map[string]int) []string {
// 	keySlice:=make([]string,0)

//		for key := range m {
//			keySlice=append(keySlice, key)
//		}
//		return keySlice
//	}
//
// -------------------------------------------------------------------------------------------------
// Упаковка значений map в строки
// Описание: Напишите функцию, которая упаковывает все значения map в одну строку, разделенную запятой.
package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"a": "s", "b": "a", "c": "l", "d": "o", "e": "m"}
	fmt.Println(a(m))
}
func a(m map[string]string) string {
	str := ""

	for _, value := range m {

		str += value + ","
	}
	return str
}
