// •	Объединить два отсортированных массива в один отсортированный.
// •	Найти длину самого длинного подмассива, в котором все элементы различны.
// •	Найти все подмассивы, сумма которых равна заданному числу.
// •	Найти пару элементов в массиве, сумма которых равна заданному числу.
// •	Найти наименьший положительный элемент, отсутствующий в массиве.
// •	Найти максимальную сумму подмассива с условием, что подмассив не должен содержать более двух различных элементов.
// •	Найти максимальную длину подмассива, сумма элементов которого равна заданному числу.
// •	Найти максимальное произведение трех чисел в массиве.
// •	Найти подмассив с максимальной суммой.
// •	Переместить все отрицательные числа в начало массива, сохраняя порядок остальных чисел.
// •	Найти подмассив с наибольшей длиной, сумма элементов которого равна нулю.
// •	Найти наибольший общий делитель всех элементов массива.

//-------------------------------------------------------------------------------------------------
//Найти минимальный элемент в массиве.
// package main

// import (
//     "fmt"
// )

// func main() {
//     arr := [7]int{10, 5, 3, 8, 7, 2, 9}
//     min := findMin(arr)
//     fmt.Println(min)
// }

// func findMin(arr [7]int) int {
//     min := arr[0]
//     for _, value := range arr {
//         if value < min {
//             min = value
//         }
//     }
//     return min
// }
//-------------------------------------------------------------------------------------------------
// Подсчитать количество положительных чисел в массиве.
// package main

// import (
//     "fmt"
// )

// func main() {
//     arr := [7]int{10, -5, -3, -8, -7, 2, 9}
//     min := positivNum(arr)
//     fmt.Println(min)
// }

// func positivNum(arr [7]int) int {
//     count := 0
//     for _, value := range arr {
//         if value > 0 {
//             count++
//         }
//     }
//     return count
// }
//-------------------------------------------------------------------------------------------------
// Найти сумму всех элементов массива.
// package main

// import (
//     "fmt"
// )

// func main() {
//     var arr [7]int=[7]int{10, 1, 0, 0, 0, 0, 0}
//     min := sum(arr)
//     fmt.Println(min)
// }

//	func sum(arr [7]int) int {
//	    sum := 0
//	    for _, value := range arr {
//	            sum+=value
//	    }
//	    return sum
//	}
//
// -------------------------------------------------------------------------------------------------
// Найти среднее значение всех элементов массива.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var arr [7]int = [7]int{15, 25, 0, 0, 0, 0, 0}
// 	min := sum(arr)
// 	fmt.Println(min)
// }

// func sum(arr [7]int) int {
// 	sum := 0
// 	count := 0

// 		for _, value := range arr {
// 			sum += value
// 			if value > 0 {
// 				count++
// 			}
// 		}
// 		return sum / count
// 	}
//
// -------------------------------------------------------------------------------------------------
// Удалить все вхождения заданного числа из массива.
// package main

// import "fmt"

// func main() {
// 	slice:=[]int{1,11,2,3,4,5,11,6,7,8,9,10,11,11,11,11,11,11,11}
// 	deletNumber(11,slice)
// }

//	func deletNumber(n int,slice []int)  {
//		buffer:=[]int{}
//		for _, v := range slice {
//			if v!=n{
//				buffer = append(buffer, v)
//			}
//		}
//		fmt.Println(buffer)
//	}
//
// -------------------------------------------------------------------------------------------------
// •	Найти все индексы заданного числа в массиве.
// package main

// import "fmt"

// func main() {
// 	massive:=[10]int{1,2,3,4,5,6,7,8,9,10}
// 	slice:=massive[:]
// 	a(10,slice)

// }
//
//	func a(n int,slice []int)  {
//		var buffer []int
//		for i, v := range slice {
//			if v==n{
//				buffer = append(buffer, i)
//			}
//		}
//		fmt.Println(buffer)
//	}
//
// -------------------------------------------------------------------------------------------------
// Создать копию массива.
// package main

// import "fmt"

//	func main() {
//		var a [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
//		b:=a[:]
//		coppyA:=coppyArray(b)
//		fmt.Println(coppyA)
//	}
//
//	func coppyArray(slice []int) []int {
//		buffer:=make([]int,len(slice))
//		copy(buffer,slice)
//		return buffer
//	}
//
// -------------------------------------------------------------------------------------------------
// Объединить два массива.
// package main

// import "fmt"

//	func main() {
//		a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//		b := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
//		sliceA := a[:]
//		sliceB := b[:]
//		mergedSlice := mergeArray(sliceA, sliceB)
//		fmt.Println(mergedSlice)
//	}
//
//	func mergeArray(a, b []int) []int {
//		c:=a[:]
//		c = append(c, b...)
//		return c
//	}
//
// -------------------------------------------------------------------------------------------------
// Поменять местами максимальный и минимальный элементы массива.
// package main

// import "fmt"

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	swap(a)
// }
// func swap(a []int) {
// 	maxIndex := 0
// 	minIndex := 0

// 	for i, v := range a {
// 		if v > a[maxIndex] {
// 			maxIndex = i
// 		}
// 		if v < a[minIndex] {
// 			minIndex = i
// 		}

//		}
//		a[maxIndex], a[minIndex] = a[minIndex], a[maxIndex]
//		fmt.Println(a)
//	}
//
// -------------------------------------------------------------------------------------------------
// Проверить, является ли массив палиндромом.
// package main

// import "fmt"

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,10,9,8,7,6,5,4,3,2,1}
// 	fmt.Println(IsPalindrom(a))
// }
// func IsPalindrom(a []int) bool {
// 	left := 0
// 	right := len(a) - 1
// 	for left < right {
// 		if a[left] != a[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}

//		return true
//	}
//
// -------------------------------------------------------------------------------------------------
// Найти второе наибольшее число в массиве.
// package main

// import "fmt"

// func main() {
// 	slice := []int{64, 34, 25, 12, 22, 11, 90}
// 	bubbleSort(slice)
// 	fmt.Println(slice)
// 	fmt.Println(slice[5])
// }

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
// -------------------------------------------------------------------------------------------------
// Перевернуть массив.
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}
// 	reverse(slice)
// 	fmt.Println(slice)

// }
//
//	func reverse(slice []int) {
//		left := 0
//		right := len(slice) - 1
//		for left < right {
//			slice[left], slice[right] = slice[right], slice[left]
//			left++
//			right--
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Удалить дубликаты из массива.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	slice := []int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 9, 10}
// 	deleteDublicate(slice)
// 	// fmt.Println(∆)
// }
// func deleteDublicate(slice []int) {
// 	newSlice := []int{}
// 	for i := 0; i < len(slice); i++ {
// 		flag := false
// 		for j := 0; j < len(newSlice); j++ {
// 			if slice[i] == newSlice[j] {
// 				flag = true
// 				break
// 			}
// 		}
// 		if !flag {
// 			newSlice = append(newSlice, slice[i])
// 		}
// 	}
// 	fmt.Println(newSlice)

// }
// -------------------------------------------------------------------------------------------------
// Переместить все нули в конце массива, сохраняя порядок ненулевых элементов.
// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 0, 2, 3, 4, 5, 6, 0, 7, 8, 9, 0, 10, 11, 12, 0, 13, 14}
// 	nulToEnd(slice)
// }

//	func nulToEnd(slice []int) {
//		numbers := []int{}
//		zero := []int{}
//		for _, v := range slice {
//			if v == 0 {
//				zero = append(zero, v)
//			} else {
//				numbers = append(numbers, v)
//			}
//		}
//		numbers = append(numbers, zero...)
//		fmt.Println(numbers)
//	}
//
//-------------------------------------------------------------------------------------------------
// // Найти пересечение двух массивов.
// package main

// import "fmt"

// func main() {
// 	first := []int{1, 1,8,1, 2, 3, 4}
// 	second := []int{1, 1, 8, 8, 4}
// 	a(first, second)
// }
// func a(first, second []int) {
// 	result := []int{}

// 	for i := 0; i < len(first); i++ {
// 		exist := false

// 		for _, v := range result {
// 			if v==first[i]{
// 				exist=true
// 				break
// 			}
// 		}
		
// 		if exist {
// 			continue
// 		}

// 		for j := 0; j < len(second); j++ {
// 				if first[i] == second[j] {
// 					result = append(result, first[i])
// 					break
// 				}
// 			}
// 		}
// 		fmt.Println(result)

// 	}
//-------------------------------------------------------------------------------------------------
// Проверить, является ли массив подмножеством другого массива.
package main

func main() {
	
}
