package main

import "fmt"

func main() {
	//Array
	var nos [5]int
	for i := 0; i < 5; i++ {
		nos[i] = i
	}
	fmt.Println("nos length = ", len(nos))
	fmt.Println(nos)

	//initilization while creation
	//var products [5]string = [5]string{"Pen", "Pencil", "Marker", "Scribble Pad", "Mouse"}
	var products = [5]string{"Pen", "Pencil", "Marker", "Scribble Pad", "Mouse"}
	fmt.Printf("Type of products = %T\n", products)
	fmt.Printf("No of products = %v\n", len(products))

	//Iterating using range
	for _, value := range products {
		fmt.Println(value)
	}

	//multidimenstional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row-%d*col-%d", i, j)
		}
	}
	fmt.Printf("Matrix - %v\n", matrix)

	//Slices
	var nosSlice []int = []int{9, 4, 2, 6, 1}
	fmt.Println("nosSlice = ", nosSlice)

	//adding new items to the slice using the 'append' function
	nosSlice = append(nosSlice, 5, 14, 12)
	fmt.Println("After appending, nosSlice = ", nosSlice)

	anotherNosSlice := []int{100, 200, 300}
	nosSlice = append(nosSlice, anotherNosSlice...)
	fmt.Println("After appending, nosSlice = ", nosSlice)

	//slicing a Slice
	/*
		nosSlice[lo:hi] => [0:4] = [0, 1, 2, 3]
		nosSlice[lo:lo] => empty
		nosSlice[lo:lo+1] => item at 'lo'
		nosSlice[lo:] => all the values from 'lo'
		nosSlice[:hi] => all the values from 0 to 'hi-1'
	*/
	fmt.Printf("nosSlice[1:4] = %v\n", nosSlice[1:4])
	fmt.Printf("nosSlice[1:1] = %v\n", nosSlice[1:1])
	fmt.Printf("nosSlice[1:2] = %v\n", nosSlice[1:2])
	fmt.Printf("nosSlice[5:] = %v\n", nosSlice[5:])
	fmt.Printf("nosSlice[:5] = %v\n", nosSlice[:5])

}
