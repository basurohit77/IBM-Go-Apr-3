package main

import "fmt"

func main() {
	operations := map[int]func(int, int) int{
		1: add,
		2: subtract,
		3: multiply,
		4: divide,
	}
	for {
		choice := getMenuOption()
		operation, exists := operations[choice]
		if !exists {
			break
		}
		n1, n2 := getOperands()
		result := operation(n1, n2)
		fmt.Printf("Result = %d\n", result)
	}
}

func log(n1, n2 int, operation func(int, int) int) int {
	fmt.Printf("Processing %d and %d\n", n1, n2)
	return operation(n1, n2)
}

func add(n1, n2 int) int {
	return n1 + n2
}

func subtract(n1, n2 int) int {
	return n1 - n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}

func divide(n1, n2 int) int {
	return n1 / n2
}

func getOperands() (n1 int, n2 int) {
	fmt.Println("Enter Operand 1")
	fmt.Scanf("%d", &n1)
	fmt.Println("Enter Operand 2")
	fmt.Scanf("%d", &n2)
	return
}

func getMenuOption() int {
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanf("%d", &choice)
	return choice
}
