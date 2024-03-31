package main

import (
	"fmt"
	"math"
)

func main(){
	var r int
	fmt.Print("masukan jari-jari: ")
	fmt.Scanln(&r)

	var luas float64 = (22.0/7) * math.Pow(float64(r), 2)
	fmt.Printf("Luas lingkaran( r = %d ) adalah %g", r, luas)
}