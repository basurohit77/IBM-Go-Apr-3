package main

import "fmt"

func main() {
	defer onMainExit("main is exiting")
	fmt.Println("main executed")
}

func onMainExit(str string) {
	fmt.Println(str)
}
