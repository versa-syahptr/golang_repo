package main

import "fmt"

func main(){
	var a,b int
	fmt.Scan(&a, &b)
	hasil := a-b
	if hasil < 0 {
		hasil *= -1
	}
	fmt.Print("Nilai mutlak dari ",a," dikurangi ",b," adalah ",hasil)
}