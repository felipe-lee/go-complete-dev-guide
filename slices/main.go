package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenOrOdd := ""
	for _, num := range nums {
		if num%2 == 0 {
			evenOrOdd = "even"
		} else {
			evenOrOdd = "odd"
		}

		fmt.Printf("%d is %v\n", num, evenOrOdd)
	}
}
