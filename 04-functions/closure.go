package main

import "fmt"

func main() {
        fmt.Println("from main")
	count := getCounter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	main2()
}

func main2() {
	fmt.Println("from main2")
	var count = getCounter()
	fmt.Println(count()) // counter will be 7
	fmt.Println(count())
	fmt.Println(count())
}
// closure is something a outer function cantain a variable, tha t variable is process in inner function and returned
// Thus the inner function keeps the variable alive of the outer function.
func getCounter() func() int {
	var counter int = 0
	return func() int {
		counter += 1
		return counter
	}
}
