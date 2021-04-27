package main

import "fmt"

func main() {
	x := 10
	//address of x => &x
	fmt.Printf("value of x is %d\n", x)
	fmt.Printf("address of x is %v\n", &x)

	var addressOfX *int = &x
	fmt.Printf("address of x is %v\n", addressOfX)

	//accessing the value using the address is called "dereference"
	fmt.Printf("Value of x using the address of x %d \n", *addressOfX)

	n1 := 10
	n2 := 20
	fmt.Printf("Before swapping n1 = %d and n2 = %d\n", n1, n2)
	var addn1 *int = &n1
	var addn2 *int = &n2

	swap(addn1, addn2)
	fmt.Printf("After swapping n1 = %d and n2 = %d\n", n1, n2)
}

func swap(addn1 *int, addn2 *int) { // Pass by reference
    /*
	var temp = *addn1
	*addn1 = *addn2
	*addn2 = temp
	*/
	*addn1, *addn2 = *addn2, *addn1

}
