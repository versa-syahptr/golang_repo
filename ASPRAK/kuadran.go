package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x > 0 && y > 0 {
		fmt.Println("Kuadran I")
	} else if x < 0 && y > 0 {
		fmt.Println("Kuadran II")
	} else if x < 0 && y < 0 {
		fmt.Println("Kuadran III")
	} else if x > 0 && y < 0 {
		fmt.Println("Kuadran IV")
	} else {
		fmt.Println("Origin")
	}
}
