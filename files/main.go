package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Expected file name to be passed as an argument.")
		os.Exit(1)
	}

	filename := os.Args[1]

	fmt.Printf("File name input: %s\n", filename)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Unable to open the file named %s. Error: %v\n", filename, err)
		os.Exit(2)
	}

	fmt.Println("File contents:")

	io.Copy(os.Stdout, file)
}
