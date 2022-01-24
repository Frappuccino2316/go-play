package main

import "fmt"

func main() {
	fmt.Println("First Process")
	goto L
	fmt.Println("Second Process")
L:
	fmt.Println("Third Process")
}

// func pow(p *[3]int) {
// 	i := 0
// 	for i < 3 {
// 		p[i] = p[i] * p[i]
// 		i++
// 	}
// }

// func main() {
// 	p := [3]int{1, 2, 3}
// 	fmt.Println(p)
// 	pow(&p)
// 	fmt.Println(p)
// }

// func main() {
// 	i := 1
// 	fmt.Println(i)
// 	inc(&i)
// 	inc(&i)
// 	inc(&i)
// 	fmt.Println(i)
// }

// func inc(p *int) {
// 	*p++
// }

// func main() {
// 	var i int
// 	p := &i
// 	i = 3
// 	fmt.Println(*p)

// 	*p = 10
// 	fmt.Println(i)
// }

// func main() {
// 	var i int
// 	p := &i
// 	pp := &p
// 	fmt.Println(p)
// 	fmt.Printf("%T\n", p)
// 	fmt.Println(pp)
// 	fmt.Printf("%T\n", pp)
// }

// map
// func main() {
// 	m := make(map[string]int)
// 	m["one"] = 1
// 	m["two"] = 2
// 	fmt.Println(m)
// }

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
