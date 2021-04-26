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
}

/*
// Normal function
func greet(firstname string, lastname string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!", firstname, lastname)
}
*/


func getRand(limit int) (rand1 int, rand2 int) {
	rand1 = rand.Intn(limit)
	rand2 = rand.Intn(limit)
	return
}
