package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(bin(number))
}

func bin(x int) int {
	if x == 0 {
		return 0
	} else {
		return (x%2) + 10*bin(x/2)
	}
}

