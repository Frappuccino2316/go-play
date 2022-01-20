package main

import "fmt"

// map
func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
}

// スライス
// func main() {
// 	// n := make([]int, 5, 10)
// 	n := []int{1, 2, 3}
// 	n = append(n, 4, 5, 6)
// 	fmt.Println("n: ", n)
// 	fmt.Println(len(n))
// 	fmt.Println(cap(n))
// }

// func main() {
// 	f := later()
// 	fmt.Println(f("Golang"))
// 	fmt.Println(f("is"))
// 	fmt.Println(f("awesome!"))
// 	fmt.Println(f(""))
// }

// func later() func(string) string {
// 	var store string

// 	return func(next string) string {
// 		s := store
// 		store = next
// 		return s
// 	}
// }

// ラベル
// func main() {
// LOOP:
// 	for {
// 		for {
// 			for {
// 				fmt.Println("Three nest!")
// 				break LOOP
// 			}
// 		}
// 	}
// 	fmt.Println(("完了"))
// }
