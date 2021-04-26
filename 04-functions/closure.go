package main

import "fmt"

func main() {
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
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
}

func getCounter() func() int {
	var counter int = 0
	return func() int {
		counter += 1
		return counter
	}
}
