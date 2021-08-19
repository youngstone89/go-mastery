package main

import (
	"fmt"
	"os"
)

func main() {

	openFile("panic.txt")

	print("Done")
}

func openFile(s string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR",r)
		}
	}()

	f, err := os.Open(s)
	if err != nil {
		panic(err)
	defer f.Close()
	}
}
