package main

import "fmt"

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
func main() {
LOOP:
	for {
		for {
			for {
				fmt.Println("Three nest!")
				break LOOP
			}
		}
	}
	fmt.Println(("完了"))
}
