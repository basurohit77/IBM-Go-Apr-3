package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(10, 20, 30))
	fmt.Println(add(10, 20, 30, 40, 50))
}

// Dynamic number of argument is known ar variadic function
func add(nos ...int) int {
	sum := 0
	for i := 0; i < len(nos); i++ {
		sum += nos[i]
	}
	return sum
}
