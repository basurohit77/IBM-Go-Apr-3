package main

import "fmt"

func main() {
	adder := getAdder()
	fmt.Println(adder(100, 200))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(300))
}

//Functions can be returned as return values from other functions
func getAdder() func(int, int) int {
	return func(i1, i2 int) int {
		return i1 + i2
	}
}

func getAdderFor(i1 int) func(int) int {
	return func(i2 int) int {
		return i1 + i2
	}
}
