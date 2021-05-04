package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileHandle, fileError := os.Open("sample.txt")
	defer fileHandle.Close()
	if fileError != nil {
		log.Fatalln(fileError)
	}
	inputReader := bufio.NewReader(fileHandle)
	for {
		line, err := inputReader.ReadString('.')
		fmt.Println(line)
		if err == io.EOF {
			break
		}
	}
	fileHandle.Close()
}
