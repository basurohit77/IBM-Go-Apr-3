package main

import (
	"fmt"
	"sync"
	"time"
)

func print(message string, delay time.Duration, ch1 chan string, ch2 chan string, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		<-ch1
		fmt.Println(message)
		fmt.Println(i)
		time.Sleep(delay * time.Second)
		ch2 <- "done"
		if i == 4 && message == "World"{
			fmt.Println(" in world")
			//msg := <-ch2
			//fmt.Println(msg)
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)
	wg.Add(2)
	go print("Hello", 1, ch1, ch2, &wg)
	go print("World", 1, ch2, ch1, &wg)
	ch1 <- "start"

	wg.Wait()
	time.Sleep(5)
	//var input string
	//fmt.Scanln(&input)
	//fmt.Println("Job Done!")
}

/*
	for writing data into the channel
		ch <- data
	for reading data from the channel
		x := <- ch
*/
