package main

import	"fmt"
import	"io/ioutil"
import	"os"
import	"strings"


func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), " ")


	strlen := len(sliceData)

	stringRank := map[int]int{}

	for i := 0; i < strlen; i++ {

		wordLen := len(sliceData[i])
		//fmt.Println(sliceData[i], wordLen)
		stringRank[wordLen] = stringRank[wordLen] + 1
		//fmt.Println(stringRank[wordLen])

	}

	for wl, rank := range stringRank {
		fmt.Printf("wordlen = %d, Rank = %d \n", wl, rank)
	}

	// Now find the word with a length used max number od time

}
