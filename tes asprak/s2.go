package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c,n int
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a, &b, &c)
		fmt.Println(singa(a,b,c))
	}
}

func singa(x,y,z int) string {
	var j1, j2 int
	j1 = int(math.Abs(float64(x-z)))
	j2 = int(math.Abs(float64(y-z)))
	if j1 < j2 {
		return "Singa A"
	} else if j1 > j2 {
		return "Singa B"
	} else {
		return "Rusa C"
	}
}