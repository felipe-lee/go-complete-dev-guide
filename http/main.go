package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)

	numBytes, errBody := resp.Body.Read(bs)

	if errBody != nil {
		fmt.Printf("Error reading body: %v\n", err)
	}

	fmt.Println(numBytes, errBody)
	fmt.Println(string(bs))
}
