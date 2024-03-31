package main

import "fmt"

func main(){
	var a1x, a1y, b1x, b1y int
	var a2x, a2y, b2x, b2y int
	var a3x, a3y, b3x, b3y int
	var g1, g2, g3 float64

	fmt.Scan(&a1x, &a1y, &b1x, &b1y)
	fmt.Scan(&a2x, &a2y, &b2x, &b2y)
	fmt.Scan(&a3x, &a3y, &b3x, &b3y)

	g1 = float64(a1y-b1y)/float64(a1x-b1x)
	g2 = float64(a2y-b2y)/float64(a2x-b2x)
	g3 = float64(a3y-b3y)/float64(a3x-b3x)

	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(g3)
}
