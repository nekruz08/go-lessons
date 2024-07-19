// package main

// import "fmt"

// func main() {
// 	for10(10)
// }

// func for10(n int)  {
// 	if n>0{
// 		sum:=0.0
// 		for i := 1; i <= n; i++ {
// 			sum+=1.0/float64(i)
// 		}
// 		fmt.Printf("%.2f\n", sum)

//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var N int
// 	fmt.Print("Введите целое число N (> 0): ")
// 	fmt.Scan(&N)

// 	sum := 0

// 	for i := N; i <= 2*N; i++ {
// 		sum += i * i
// 	}

//		fmt.Printf("Сумма квадратов от %d^2 до %d^2 равна %d\n", N, 2*N, sum)
//	}
//
// -------------------------------------------------------------------------------------------------
// page28/series1/series6
// package main

// import "fmt"

// func main() {
// 	newSlice:=[10]float64{}

// 	for _, v := range newSlice {
// 		v+=1
// 	}
// 	fmt.Println(newSlice)

//		sum:=0.0
//		for _, v := range newSlice {
//			sum+=v
//		}
//		fmt.Println(sum)
//	}
//
// -------------------------------------------------------------------------------------------------
// series2
// package main

// import "fmt"

// func main() {
// 	series2()
// }

// func series2()  {
// 	var p float64=1
// 	var n float64
// 	for i := 0; i < 10; i++ {
// 		fmt.Scan(&n)
// 		p=p*p
// 	}
// 	fmt.Println(p)
// }
// -------------------------------------------------------------------------------------------------
// series6
package main

import "fmt"

func main() {
	series2()
}

func series2()  {
	var p float32
	var n int
	fmt.Scan(&n)

	for i := 0; i < 10; i++ {
		fmt.Scan(&p)
		fmt.Println(p-float32((int(p)))	)

	}
}
