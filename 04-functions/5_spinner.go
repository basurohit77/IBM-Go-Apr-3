package main

import "fmt"

func main() {
	increment, decrement := getSpinner()
	fmt.Println("Increment = ", increment())
	fmt.Println("Increment = ", increment())
	fmt.Println("Increment = ", increment())
	fmt.Println("Increment = ", increment())

	fmt.Println("Decrement = ", decrement())
	fmt.Println("Decrement = ", decrement())
	fmt.Println("Decrement = ", decrement())
	fmt.Println("Decrement = ", decrement())
	fmt.Println("Decrement = ", decrement())

}

// A spinner is also a closure with return of more than one inner function, both the inner function use and process the variable of outer function.
func getSpinner() (increment func() int, decrement func() int) {
	counter := 0
	increment = func() int {
		counter += 1
		return counter
	}
	decrement = func() int {
		counter -= 1
		return counter
	}
	return
}
