package main

import (
	"fmt"
	"math/rand"
)

/*
var greet func(string, string) string = func(firstname, lastname string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!", firstname, lastname)
}
*/

func main() {

	
        // Function on variable
        greet := func(firstname, lastname string) string {
		return fmt.Sprintf("Hi %s %s, Have a nice day!", firstname, lastname)
	} 

	var firstname, lastname string
	fmt.Println("Enter the first name :")
	fmt.Scanf("%s", &firstname)

	fmt.Println("Enter the last name :")
	fmt.Scanf("%s", &lastname)

	greetMsg := greet(firstname, lastname) //Use the variable
	fmt.Println(greetMsg)

	//fmt.Println(getRand(200))
	v1, v2 := getRand(1000)
	fmt.Println(v1, v2)
        
       // operation returns 5, 10 to getRand, and then getRand returns rand1, rand2 to log function
        v3, v4 := log(5, 10, getRand)
	fmt.Println(v3, v4)

}

/*
// Normal function
func greet(firstname string, lastname string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!", firstname, lastname)
}
*/

// Functions can be passed as arguments to other functions
func log(m1, m2 int, operation func(int, int) (int, int) ) (int, int) {
	fmt.Println("process m1 and m2")
	return operation(m1, m2)
}

func getRand(limit int) (rand1 int, rand2 int) {
	rand1 = rand.Intn(limit)
	rand2 = rand.Intn(limit)
	return
}
