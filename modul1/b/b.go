package main

import "fmt"

func main(){
	var luas float64
	var r int64

	fmt.Scan(&r)
	luas = (22.0/7) * float64(r)
	fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luas)
}