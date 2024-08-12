// package main

// import "fmt"

//	func main() {
//		var arr [5]int
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// Создание массива
// - Пустой массив:
// package main

// import "fmt"

//	func main() {
//		var arr [5]int
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// - Инициализация массива с начальными значениями:
// package main

// import "fmt"

//	func main() {
//		var arr [5]int=[5]int{1,2,3,4,5}
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		arr:=[5]int{1,2,3,4,5}
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
//- Автоматическое определение длины массива:
// package main

// import "fmt"

//	func main() {
//		arr:=[...]int{1,2,3,4,5,6,7,8,9,10}
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		var matrix [3][4]int
//		matrix[0][0] = 1
//		matrix[0][1] = 2
//		matrix[1][0] = 1
//		matrix[1][1] = 2
//		matrix[2][3] = 10
//		fmt.Println(matrix)
//	}
//
// -------------------------------------------------------------------------------------------------
// - Ключи (индексы) массивов
// package main

// import "fmt"

//	func main() {
//		arr:=[5]int{0,0,0,0,0}
//		fmt.Println(arr)
//		arr[4]=1
//		fmt.Println(arr)
//		// fmt.Println(arr[5])
//	}
//
// -------------------------------------------------------------------------------------------------
// - Получение и изменение элементов по индексу
// package main

// import "fmt"

//	func main() {
//		arr:=[5]int{1,2,3,4,5}
//		fmt.Println(arr)
//		fmt.Println(arr[4])
//		arr[3]=3
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// - Длина (len) и вместимость (cap) массива
// package main

// import "fmt"

//	func main() {
//		var arr [5]int
//		fmt.Println(len(arr))
//		fmt.Println(cap(arr))
//	}
//
// -------------------------------------------------------------------------------------------------
// - Итерация по массиву
// package main

// import "fmt"

// func main() {
// 	arr:=[10]int{1,2,3,4,5,6,7,8,9,10}
// 	// fmt.Println(arr)

//		for i := 0; i < len(arr); i++ {
//			fmt.Printf("%v ",arr[i])
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var arr = [...]int{1, 2, 4, 5, 6, 7, 8,9,10}

//		for _, v := range arr {
//			fmt.Printf("%v ",v)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Что такое срез
// package main

// import "fmt"

//	func main() {
//		var slice []int
//		var newSlice []int=[]int{1,2,3,4,5,6,7}
//		slice = append(slice, newSlice...)
//		fmt.Println(slice)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	s := make([]int, 0)
// 	fmt.Println(s, len(s), cap(s))
// 	s = append(s, 1)
// 	fmt.Println(s, len(s), cap(s))

//		s = make([]int, 0, 10)
//		fmt.Println(s, len(s), cap(s))
//		s = append(s, 1)
//		fmt.Println(s, len(s), cap(s))
//	}
//
// -------------------------------------------------------------------------------------------------
// Создание слайса
// Срезы можно создавать несколькими способами:
// - На основе массива:
// package main

// import "fmt"

// func main() {
// 	arr:=[5]int{1,2,3,4,5}
// 	slice:=arr[1:4] // срез содержит элементы со 2-го по 4-й (индексы 1, 2, 3)
// 	fmt.Println(slice)
// }
// -------------------------------------------------------------------------------------------------
// - С помощью литерала:
// package main

// import "fmt"

// func main() {
// 	slice:=[]int{1,2,3,4,5,6,7,8,9,10}
// 	fmt.Println(slice)
// }
// -------------------------------------------------------------------------------------------------
// - Пустой срез:
// package main

// import "fmt"

// func main() {
// var emptySlice []int
// fmt.Println(emptySlice==nil)
// fmt.Println(len(emptySlice))
// fmt.Println(cap(emptySlice))
// fmt.Println(emptySlice)

// }
// -------------------------------------------------------------------------------------------------
// Длина и вместимость слайса
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))

// }
// -------------------------------------------------------------------------------------------------
// Создание слайса через make
// package main

// import "fmt"

//	func main() {
//		slice := make([]int, 5,10) //первым элементом мы инициализируем столько то нулей вторым параметром зададим сколько элементов может вместит даанный слайс)
//		fmt.Println(slice)
//		fmt.Println(len(slice))
//		fmt.Println(cap(slice))
//	}
//
// -------------------------------------------------------------------------------------------------
// Функция append
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	slices := []int{1, 2, 3}
// 	fmt.Println(len(slices))
// 	fmt.Println(cap(slices))

// 	slices = append(slices, 4, 5, 6, 7, 8, 9, 10)
// 	fmt.Println(len(slices))
// 	fmt.Println(cap(slices))
// 	slices=append(slices, 11)
// 	fmt.Println(slices)

// }
// -------------------------------------------------------------------------------------------------
// Конкатенация срезов
// package main

// import "fmt"

//	func main() {
//		slice1:=[]int{1,2,3}
//		slice2:=[]int{4,5}
//		slice3:=append(slice1,slice2...)
//		fmt.Println(slice3)
//	}
//
// -------------------------------------------------------------------------------------------------
// Операция слайсинг (создание среза на основе массива или слайса)
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slices := arr[1:4] // 234
// 	fmt.Println(slices)

//		slices2:=slices[3:4]
//		fmt.Println(slices2)
//	}
//
// -------------------------------------------------------------------------------------------------
// Функция copy
// package main

// import "fmt"

// func main() {
// 	src := []int{1, 2, 3}
// 	var dst []int {}

// 	copy(dst, src)
// 	fmt.Println(dst)
// }
// var nilSlice []int

// -------------------------------------------------------------------------------------------------
// Пустой слайс
// package main

// import "fmt"

// func main() {
// 	var slice []int = []int{}
// 	fmt.Println(slice==nil)
// 	fmt.Println(len(slice))
// 	fmt.Println(cap(slice))
// }
// -------------------------------------------------------------------------------------------------
// nil слайс
// package main

// import "fmt"

//	func main() {
//		var slice []int
//		fmt.Println(slice==nil)
//		// после инициализации уже не nil
//		slice=[]int{}
//		fmt.Println(slice==nil)
//		fmt.Println(len(slice))
//		fmt.Println(cap(slice))
//	}
//
// -------------------------------------------------------------------------------------------------
// Инициализация массива с использованием ключей (индексов)
// package main

// import "fmt"

//	func main() {
//		arr:=[5]int{0:10,4: 20}
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// Копирование массива
// В отличие от срезов, при присваивании одного массива другому происходит
// копирование значений:
// Это значить если поменять второй  массив базовый т.е. первый не меняетсся, т.е. происходить реальное копированеи
// тогда как в массиве по мимо копирование она хранить в себе указатель на базовый массив изиеняя слаййс мы
// изменяем базоввый массив
// с копированием массивово такое не прокатит
// с копийе что не поделай базовый массив остаетс неизменныи

// package main

// import "fmt"

// func main() {
// 	arr1:=[]int{1,2,3}
// 	arr2:=arr1 // копирование значений
// 	arr2[0]=10
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)

// }
// -------------------------------------------------------------------------------------------------
// Дополнительные аспекты срезов
// package main

// import "fmt"

//	func main() {
//		slice:=[]int{1,2,3,4,5}
//		slice2:=slice[1:3]
//		fmt.Println(len(slice2),cap(slice2))
//	}
//
// -------------------------------------------------------------------------------------------------
// Изменение среза изменяет исходный массив
// package main

// import "fmt"

// func main() {
// 	arr:=[5]int{1,2,3,4,5}
// 	slice:=arr[1:4]
// 	// fmt.Println(slice)

//		slice[0]=10
//		fmt.Println(arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// Nil срезы
// package main

// import "fmt"

//	func main() {
//		var nilSlice []int
//		fmt.Println(nilSlice==nil)
//		fmt.Println(len(nilSlice),cap(nilSlice))
//	}
//
// -------------------------------------------------------------------------------------------------
// Срез с максимальной вместимостью
// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	slice := arr[1:3:4] // длина 2, вместимость 3
// 	fmt.Println(slice)
// 	fmt.Println(len(slice), cap(slice))

// }
// -------------------------------------------------------------------------------------------------
// Удаление элемента из среза
// package main

// import "fmt"

//	func main() {
//		slice:=[]int{1,2,3,4,5}
//		index:=2
//		slice=append(slice[:index],slice[index+1:]... )
//		fmt.Println(slice)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		slice:=[]int{1,2,3,4,5}
//		index:=2
//		slice = append(slice[:index],slice[index+1:]... )
//		fmt.Println(slice)
//	}
//
// -------------------------------------------------------------------------------------------------
// Вставка элемента в срез
// Для вставки элемента в срез требуется создание нового среза:
// package main

// import "fmt"

// func main() {
// slice:=[]int{1,2,3,4,5}
// index:=3 // индекс для вставки
// value:=4 // значение для вставки
// slice = append(slice[:index],append([]int{value},slice[index:]...)... )
// fmt.Println(slice)
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		slice:=[]int{1,2,3,4,5,7}
//		index:=5
//		value:=6
//		slice = append(slice[:5], append([]int{value},slice[index:]...)...)
//		fmt.Println(slice)
//	}
//
// -------------------------------------------------------------------------------------------------
// Общая формула вычисления ёмкости среза
// package main

// import "fmt"

// func main() {
// 	arr:=[20]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
// 	fmt.Println(arr)
// 	slice:=arr[5:10:15]
// 	fmt.Println(slice)
// 	fmt.Println(len(slice),cap(slice))
// }

// -------------------------------------------------------------------------------------------------
// Если используется синтаксис с двумя параметрами arr[low:high]
// то ёмкость вычисляется как: capacity=len(arr)−low
// package main

// import "fmt"

//	func main() {
//		var arr [20]int=[20]int{1,2,3,4,5,6,7,8,9,10}
//		slice:=arr[5:10]
//		fmt.Println(slice)
//		fmt.Println(len(slice),cap(slice))
//	}
//
// -------------------------------------------------------------------------------------------------
// Общая формула вычисления ёмкости для пустого среза
// Если пустой срез создается с использованием функции make с заданной ёмкостью, то ёмкость
// среза равна указанному значению.
// Если ёмкость не указана, то по умолчанию она равна 0.
// package main

// import "fmt"

//	func main() {
//		slice:=make([]int,0,10)
//		fmt.Println(slice==nil)
//		fmt.Println(len(slice),cap(slice))
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var nilSlice []int
// 	fmt.Println(len(nilSlice),cap(nilSlice))
// 	fmt.Println(nilSlice == nil)

// 	fmt.Println()

// 	emptySlice:=make([]int,10)// при создании слайсай если не задавать  cap, то по умолчанию  cap будет равнятся len т.у(cap=len)
// 	fmt.Println(len(emptySlice),cap(emptySlice))
// 	fmt.Println(emptySlice == nil)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var nilSlice []int
// 	var emptySlice []int=[]int{}

//		fmt.Println(len(nilSlice),cap(nilSlice))
//		fmt.Println(len(emptySlice),cap(emptySlice))
//	}
//
// -------------------------------------------------------------------------------------------------
// Алгоритм увеличения ёмкости
// Если ёмкость среза была 0, то она увеличивается до 1.
// package main

// import "fmt"

//	func main() {
//		slice:=[]int{1,2,3,4}
//		fmt.Println(cap(slice)) // cap=4
//		slice = append(slice,5) // cap=8
//		fmt.Println(cap(slice))
//	}
//
// -------------------------------------------------------------------------------------------------
// провери данное утверждение:
// так продолжается до до того как размер cap слайса не равняется 1024
// package main

// import "fmt"

//	func main() {
//		slice := make([]int, 1024)
//		fmt.Println(cap(slice))
//		slice = append(slice, 1)
//		fmt.Println(cap(slice))
//	}
//
// -------------------------------------------------------------------------------------------------
// Если ёмкость среза была 0, то она увеличивается до 1.
// package main

// import "fmt"

// func main() {
// 	// slice:=make([]int,0)
// 	// fmt.Println(cap(slice))
// 	// slice = append(slice, 1)
// 	// fmt.Println(cap(slice))
// 	// slice = append(slice, 2)
// 	// fmt.Println(cap(slice))

// 	var nilSlice []int
// 	fmt.Println(nilSlice)
// 	fmt.Println(cap(nilSlice))		//0

// 	nilSlice = append(nilSlice, 1)	//1
// 	fmt.Println(nilSlice)
// 	fmt.Println(cap(nilSlice))

// 	nilSlice = append(nilSlice, 2)	// cap равно 1 и воторой элемент не вместится поэтуму 1*2=2 т.е текущий cap равняется 3
// 	fmt.Println(nilSlice)
// 	fmt.Println(cap(nilSlice))

// 	nilSlice = append(nilSlice, 3)	// прежний cap был 2 поэтому удваиваетяс, т.е будет 4
// 	fmt.Println(nilSlice)
// 	fmt.Println(cap(nilSlice))

// 	nilSlice = append(nilSlice, 4)	// прежний cap = 4 поэтому увеличить надобности нет
// 	fmt.Println(nilSlice)
// 	fmt.Println(cap(nilSlice))

// 	nilSlice = append(nilSlice, 5)	//прежний сap=4 м и пятый не вместится поэтому будеть удваиваться т.е 4*2=8
// 	fmt.Println(nilSlice)
// 	fmt.Println(cap(nilSlice))
// }

// -------------------------------------------------------------------------------------------------
// Сортировка массива и среза в Go
// Сортировка пузырьком (Bubble Sort)
// package main

// import "fmt"

//	func bubbleSort(slice []int) {
//		n := len(slice)
//		for i := 0; i < n-1; i++ {
//			for j := 0; j < n-i-1; j++ {
//				if slice[j] > slice[j+1] {
//					slice[j], slice[j+1] = slice[j+1], slice[j]
//				}
//			}
//		}
//	}
//
//	func main() {
//		slice := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
//		bubbleSort(slice)
//		fmt.Println("Bubble Sort:", slice) // [11 12 22 25 34 64 90]
//	}
//
// -------------------------------------------------------------------------------------------------
// Отсортировать срез целых чисел от 1 до 5 в порядке убывания.
// package main

// import "fmt"

//	func main() {
//		slice:=[]int{7,1,8,9,6,2,0,3,4,5}
//		bubbleSort(slice)
//		fmt.Println(slice)
//	}
//
//	func bubbleSort(slice []int)  {
//		n:=len(slice)
//		for i := 0; i < n-1; i++ {
//			for j := 0; j < n-i-1; j++ {
//				if slice[j]<slice[j+1]{
//					slice[j],slice[j+1]=slice[j+1],slice[j]
//				}
//			}
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Отсортировать срез строк, содержащих имена животных, в порядке возрастания.
// package main

// import "fmt"

// func main() {
// 	slice:=[]string{"horse","fox","lion"}
// 	fmt.Println(slice)

//		bubbleSort(slice)
//		fmt.Println(slice)
//	}
//
//	func bubbleSort(slice []string)  {
//		n:=len(slice)
//		for i := 0; i < n-1; i++ {
//			for j := 0; j < n-i-1; j++ {
//				if len(slice[j])>len(slice[j+1]){
//					slice[j],slice[j+1]=slice[j+1],slice[j]
//				}
//			}
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Отсортировать срез вещественных чисел от 0.1 до 0.5 в порядке убывания.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
// 	bubbleSort(slice)

// }
// func bubbleSort(slice []float64) {
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if slice[j] < slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// 	fmt.Println(slice)

// }
// -------------------------------------------------------------------------------------------------
// Реализовать пузырьковую сортировку для среза структур Person,
// отсортировать по возрасту в порядке убывания.
// package main

// import "fmt"

// type Person struct {
// 	Age int
// }

// func main() {
// 	slice := []Person{
// 		{
// 			Age: 17,
// 		},
// 		{
// 			Age: 12,
// 		},
// 		{
// 			Age: 15,
// 		},
// 		{
// 			Age: 8,
// 		},
// 		{
// 			Age: 13,
// 		},
// 		{
// 			Age: 16,
// 		},
// 		{
// 			Age: 9,
// 		},
// 		{
// 			Age: 14,
// 		},
// 		{
// 			Age: 11,
// 		},
// 		{
// 			Age: 1,
// 		},
// 		{
// 			Age: 10,
// 		},
// 		{
// 			Age: 7,
// 		},
// 		{
// 			Age: 6,
// 		},
// 		{
// 			Age: 5,
// 		},
// 		{
// 			Age: 4,
// 		},
// 		{
// 			Age: 3,
// 		},
// 		{
// 			Age: 2,
// 		},
// 	}
// 	bubbleSort(slice)

// }
// func bubbleSort(slice []Person) {
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if slice[j].Age < slice[j+1].Age {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// 	for _, v := range slice {
// 		fmt.Print(v.Age," ")
// 	}
// }
// -------------------------------------------------------------------------------------------------
// Сортировать срез строк по длине строки в порядке возрастания с использованием пузырьковой сортировки.
// package main

// import "fmt"

// type Person struct {
// 	Age int
// }

// func main() {
// 	slice := []string{"aaaaaaaa","aaaaaaa","aaaaaa","aaaaa","aaaa","aaa","a"}
// 	bubbleSort(slice)

// }
// func bubbleSort(slice []string) {
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if len(slice[j]) > len(slice[j+1]) {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// 	fmt.Println(slice)
// }
// -------------------------------------------------------------------------------------------------
// Отсортировать срез целых чисел с использованием пузырьковой сортировки
// и определить, если количество перестановок элементов превышает заданный порог.
// package main

// import "fmt"

// type Person struct {
// 	Age int
// }

// func main() {
// 	slice := []string{"aaaaaaaa","aaaaaaa","aaaaaa","aaaaa","aaaa","aaa","a"}
// 	bubbleSort(slice)

// }
//
//	func bubbleSort(slice []string) {
//		n := len(slice)
//		for i := 0; i < n-1; i++ {
//			for j := 0; j < n-i-1; j++ {
//				if len(slice[j]) > len(slice[j+1]) {
//					slice[j], slice[j+1] = slice[j+1], slice[j]
//				}
//			}
//		}
//		fmt.Println(slice)
//	}
//
// -------------------------------------------------------------------------------------------------
// Быстрая сортировка (Quick Sort)
// package main

// import (
// 	"fmt"
// )

// // Функция быстрой сортировки, сортирующая массив на месте
// func quickSort(arr []int, low, high int) {
// 	if low < high {
// 		// Разделяем массив и получаем индекс опорного элемента
// 		p := partition(arr, low, high)
// 		// Рекурсивно сортируем элементы до и после опорного элемента
// 		quickSort(arr, low, p-1)
// 		quickSort(arr, p+1, high)
// 	}
// }

// // Функция разделения массива на две части
// func partition(arr []int, low, high int) int {
// 	pivot := arr[high] // Опорный элемент
// 	i := low - 1       // Индекс меньшего элемента

// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i] // Меняем местами
// 		}
// 	}
// 	arr[i+1], arr[high] = arr[high], arr[i+1] // Помещаем опорный элемент на его правильное место
// 	return i + 1
// }

//	func main() {
//		arr := []int{3, 6, 8, 10, 1, 2, 1}
//		fmt.Println("Исходный массив:", arr)
//		quickSort(arr, 0, len(arr)-1)
//		fmt.Println("Отсортированный массив:", arr)
//	}
//
// -------------------------------------------------------------------------------------------------
// Поиск элементов
// package main

// import "fmt"

// func main() {
// 	slice := []int{10, 20, 30, 40, 50}
// 	index := linearSearch(slice, 30)
// 	fmt.Println("Index of 30:", index) // Index of 30: 2
// }

// func linearSearch(slice []int, target int) int {
// 	for i, v := range slice {
// 		if v == target {
// 			return i
// 		}
// 	}
// 	return -1 // элемент не найден
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// // BinarySearch выполняет бинарный поиск в отсортированном слайсе и возвращает индекс элемента или -1, если элемент не найден.
// func BinarySearch(arr []int, target int) int {
//     low, high := 0, len(arr)-1

//     for low <= high {
//         mid := low + (high-low)/2

//         if arr[mid] == target {
// 			return mid
//         } else if arr[mid] < target {
//             low = mid + 1
//         } else {
//             high = mid - 1
//         }
//     }

//     return -1
// }

// func main() {
//     arr := []int{1, 3, 5, 7, 9, 11, 13}
//     target := 7

//	    result := BinarySearch(arr, target)
//	    if result != -1 {
//	        fmt.Printf("Элемент %d найден в индексе %d\n", target, result)
//	    } else {
//	        fmt.Println("Элемент не найден")
//	    }
//	}
//
// ------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		slice := []int{4, 2, 3, 1}
//		sort.Ints(slice)
//		fmt.Println(slice) // [1 2 3 4]
//	}
//
// ------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		slice := []float64{4.1, 2.2, 3.3, 1.4}
//		sort.Float64s(slice)
//		fmt.Println(slice) // [1.4 2.2 3.3 4.1]
//	}
//
// ------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		slice := []string{"dog", "cat", "bird"}
//		sort.Strings(slice)
//		fmt.Println(slice) // [bird cat dog]
//	}
//
// ------------------------------------------------------------------------------------------------
// - sort.Sort(sort.Interface)
// Позволяет сортировать срезы пользовательских типов,
// реализующих интерфейс sort.Interface. Интерфейс требует реализации трех методов: Len(),
// Less(i, j int) bool, и Swap(i, j int).
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type Person struct {
// 	Name string
// 	Age  int
// }
// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }
// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }
// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

//	func main() {
//		people := []Person{
//			{"Alice", 3},
//			{"Bob", 2},
//			{"Charlie", 1}}
//		sort.Sort(ByAge(people))
//		fmt.Println(people) // [{Bob 25} {Alice 30} {Charlie 35}]
//	}
//
// ------------------------------------------------------------------------------------------------
// Встроенные возможности Go для поиска
// - sort.SearchInts(slice []int, x int) int
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	index := sort.SearchInts(slice, 3)
// 	fmt.Println(index) // 2

// }
// ------------------------------------------------------------------------------------------------
// Ищет значение x в отсортированном срезе вещественных чисел.
// package main

// import (
// 	"fmt"
// 	"sort"
// )

//	func main() {
//		slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
//		index := sort.SearchFloat64s(slice, 3.3)
//		fmt.Println(index) // 2
//	}
//
// ------------------------------------------------------------------------------------------------
// Ищет строку x в отсортированном срезе строк.
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"apple", "banana", "cherry"}
	index := sort.SearchStrings(slice, "banana")
	fmt.Println(index) // 1

}
