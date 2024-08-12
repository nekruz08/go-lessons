// package main

// import "fmt"

//	func main() {
//		var m map[string]int=map[string]int{"nekruz":1}
//		fmt.Println(m)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var m map[string]int
// 	fmt.Println(m==nil)

// 	fmt.Println(m["nonexistent"])

// }

// -------------------------------------------------------------------------------------------------
// package main

// func main() {
// 	var m map[string]int
// 	m["nonexistent"] = 10

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var m map[string]int
// 	m=make(map[string]int)
// 	m["nonexistent"] = 10
// 	fmt.Println(m)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m:=make(map[string]int)
// 	m["apple"]=5
// 	m["banana"]=10

// 	fmt.Println(len(m))

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m:=make(map[string]int,5)
// 	fmt.Println(m)
// 	fmt.Println(len(m))
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m:=make(map[string]int)
// 	m["key1"]=10
// 	m["key2"]=20
// 	fmt.Println(m)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m:=make(map[string]int,10)
// 	m["key1"]=1
// 	m["key2"]=2
// 	fmt.Println(m)

// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		m:=map[string]int{"a":1,"b":2,"c":3}
//		fmt.Println(m)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		m := new(map[string]int)
//		*m = make(map[string]int)
//		(*m)["key1"]=10
//		fmt.Println(*m)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var a map[string]int
// 	b := make(map[string]int)
// 	fmt.Println("a==nil:", a == nil)
// 	fmt.Println("b==nil:", b == nil)
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		m:=make(map[string]int)
//		m["a"]=1
//		m["a"]=2
//		fmt.Println(m["a"])
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m := map[string]int{
// 		"key1": 10,
// 		"key2": 20,
// 	}
// 	// value1 := m["key1"]
// 	// value2 := m["key2"]
// 	// fmt.Println(value1,value2)

//		value, exists := m["key3"]
//		if exists {
//			fmt.Println("key1 exists with value", value)
//		} else {
//			fmt.Println("key1 does not exist")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m := map[string]int{	//эсли значение непреривный то их выводить по порядку: 345
// 		"key1": 3,
// 		"key2": 4,
// 		"key3": 5,
// 	}
// 	for key, value := range m {
// 		fmt.Println(key, value)
// 	}
// }

// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
//     m := map[string]int{ //если значение не являются неприривным тогда вывод будеть случайным
//         "apple":  5,
//         "banana": 3,
//         "cherry": 7,
//     }

//	    for key, value := range m {
//	        fmt.Println(key, value)
//	    }
//	}
//
// -------------------------------------------------------------------------------------------------
// Использование map как set (фильтр уникальности)
// map можно использовать как set, когда значения не важны, а важны только уникальные ключи.
// package main

// import "fmt"

// func main() {
// 	set := make(map[string]struct{})
// 	set["item1"] = struct{}{}
// 	set["item2"] = struct{}{}

// 	if _, exists := set["item1"]; exists {
// 		fmt.Println("item1 exists")
// 	}

//		if _, exists := set["item3"]; !exists {
//			fmt.Println("item2 not exists")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Поиск O(1) вместо O(n)
// package main

// import "fmt"

//	func main() {
//		m := map[string]int{
//			"key1": 10,
//			"key2": 20,
//		}
//		_, exists := m["key1"] // поиск выполняется за O(1)
//		fmt.Println(exists)    // true
//	}
//
// -------------------------------------------------------------------------------------------------
// map можно использовать как set, когда значения не важны, а важны только уникальные
// ключи. Значением можно использовать пустую структуру struct{}, которая занимает
// нулевое количество памяти.
// package main

// import "fmt"

// func main() {
// 	set := make(map[string]struct{})
// 	set["item1"] = struct{}{}
// 	set["item2"] = struct{}{}

// 	if _, exists := set["item1"]; exists {
// 		fmt.Println("item1 exists")
// 	}
// }
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		set:=map[string]int{
//			"a":1,
//			"b":2,
//			"c":3,
//			"d":4}
//		copy:=map[string]int{}
//		for k, v := range set {
//			copy[k]=v
//		}
//		for k, v := range copy {
//			fmt.Printf("%s:%d\n",k,v)
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		m := map[string]int{
//			"key1": 10,
//			"key2": 20,
//			"key3": 30,
//		}
//		keysToDelete := []string{}
//		for key := range m {
//			keysToDelete = append(keysToDelete, key)
//		}
//		for _, key := range keysToDelete {
//			delete(m, key)
//		}
//		fmt.Println(m) // map[]
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	// m := make(map[int]string)
// 	// m[1] = "one"
// 	// m[2] = "two"

// 	type Person struct {
// 		Name string
// 		Age  int
// 	}

//		peopleMap := make(map[Person]string)
//		p := Person{Name: "Alice", Age: 30}
//		peopleMap[p] = "Programmer"
//		fmt.Println(peopleMap[p]) // Programmer
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	m := map[string]int{
// 		"key1": 10,
// 		"key2": 20,
// 	}
// 	delete(m, "key1")
// 	fmt.Println(len(m))

// 	// Для уменьшения размера map можно использовать reallocation:
// 	newMap := make(map[string]int, len(m))
// 	for k, v := range m {
// 		newMap[k] = v
// 	}
// 	m = newMap
// 	fmt.Println(len(m))

// }
// -------------------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
	type User struct {
		Name string
		Age  int
	}

	m := make(map[string]*User)
	m["user1"] = &User{Name: "Alice", Age: 30}
	m["user2"] = &User{Name: "Bob", Age: 25}

	// Изменение значения через указатель
	m["user1"].Age = 31
	fmt.Println(m["user1"].Age) // 31
}
