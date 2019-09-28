package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange", "Plum", "Pomegranate", "Grape"}
	// some := make([]string, 2)
	// copy(some, fruits[1:3])
	// some = append(some, "Banana")

	some := fruits[1:3]
	fmt.Println("s", len(some), cap(some))
	some = append(some, "Banana")

	fmt.Println("Fruits actual: ", fruits)
	fmt.Println("Some actual: ", some)

	fmt.Println("f", len(fruits), cap(fruits))
	fmt.Println("s", len(some), cap(some))

	fmt.Printf("fruits: %p\n", &fruits)
	fmt.Println()
	fmt.Printf("fruits[1]: %p\n", &fruits[1])
	fmt.Printf("some[0]: %p\n", &some[0])
	fmt.Printf("fruits[2]: %p\n", &fruits[2])
	fmt.Printf("some[1]: %p\n", &some[1])
	fmt.Printf("fruits[3]: %p\n", &fruits[3])
	fmt.Printf("some[2]: %p\n", &some[2])

	// if len(s) + len(newElements) > cap(s) {
	// 	// Allocate new backing array
	// 	// copy content (s) over to new array
	// } else {
	// 	// Just resize existing slice
	// }

}
