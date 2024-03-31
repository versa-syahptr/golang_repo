package main

import "fmt"

func main(){
	var a,b int
	var kata string
	fmt.Scan(&kata, &a, &b)

	fmt.Println("Kata =", kata)
	fmt.Println("Jumlah =", a+b)
}