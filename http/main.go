package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	lw := logWriter{}

	numBytes, errCopy := io.Copy(lw, resp.Body)

	if errCopy != nil {
		fmt.Printf("Error copying: %v\n", errCopy)
		os.Exit(2)
	}

	fmt.Printf("Num bytes: %v\n", numBytes)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
