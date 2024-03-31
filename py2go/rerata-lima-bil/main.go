package main

import "fmt"

func main(){
	var a,b,c,d,e int
	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Print("Rata-rata: ", float64(a+b+c+d+e)/5)
}