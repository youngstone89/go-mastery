package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	os.Chdir("osstat")
	fmt.Println("wordir::",dir)
	file, err := os.Open("file.txt") // For read access.

	if err != nil {
		log.Fatal(err)
	}

	info, err := os.Stat("file.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("info::",info)

	data := make([]byte, 102)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
