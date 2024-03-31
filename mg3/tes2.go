package main

import "fmt"

func main(){
	var n string

	fmt.Scan(&n)

	fmt.Printf("%c\n", n[0])
	fmt.Printf("%c\n", n[1])
	fmt.Printf("%c\n", n[2])
	fmt.Printf("%c\n", n[3])

	fmt.Printf("%s\n", n[0:2])
	fmt.Printf("%s\n", n[1:3])
	fmt.Printf("%s\n", n[2:4])

	fmt.Printf("%c%c%c\n", n[0],n[1],n[2])
	fmt.Printf("%c%c%c\n", n[1],n[2],n[3])

	
}

