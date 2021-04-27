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


	//initialization while creation
	var products = [5]string{"Pen", "Pencil", "Marker", "Scribble Pad", "Mouse"}
	fmt.Printf("Type of products = %T\n", products)
	fmt.Printf("No of products = %v\n", len(products))

	//Iterating using range
	for _, value := range products {
		fmt.Println(value)
	}

	//Multidimensional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row-%d*col-%d", i, j)
		}
	}
	fmt.Printf("Matrix - %v \n", matrix)

	// 2. Slices, here unlike in array we never initialise the size
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
	fmt.Printf("nosSlice[1:4] = %v \n", nosSlice[1:4])
	fmt.Printf("nosSlice[1:1] = %v \n", nosSlice[1:1])
	fmt.Printf("nosSlice[1:2] = %v \n", nosSlice[1:2])
	fmt.Printf("nosSlice[5:] = %v \n", nosSlice[5:])
	fmt.Printf("nosSlice[:5] = %v \n", nosSlice[:5])

	// 3. Maps
	cityRanks := map[string]int{
		"Udupi":     2,
		"Bengaluru": 4,
		"Mangaluru": 1,
		"Mysuru":    3,
	}
	fmt.Printf("Type of cityRanks = %T\n", cityRanks)
	fmt.Println(cityRanks)

	fmt.Println("Adding a new key/value pair")
	cityRanks["Chennai"] = 5
	fmt.Println(cityRanks)

	for city, rank := range cityRanks {
		fmt.Printf("City = %q, Rank = %d\n", city, rank) // %q put the city in double quotes
	}

	if chennaiRank, exists := cityRanks["Chennai"]; exists { // if cityRanks["Chennai"] present then exist is true
		fmt.Println("Rank of Chennai is ", chennaiRank)
	} else {
		fmt.Println("Chennai is not ranked") // if not present then chennaiRank is zero
	}

	//remove an key/value pair
	delete(cityRanks, "Chennai")csss
	fmt.Println("After deleting Chennai")
	if chennaiRank, exists := cityRanks["Chennai"]; exists {
		fmt.Println("Rank of Chennai is ", chennaiRank)
	} else {
		fmt.Println("Chennai is not ranked")
	}


}
