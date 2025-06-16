package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// writes files
	f, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("write data to fileeeeee\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote", size, "bytes to file.")

	f.Close()

	// reads files
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	//reads files line by line

	file2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
