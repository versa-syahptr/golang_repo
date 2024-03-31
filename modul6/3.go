package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	var dalamL1, dalamL2 bool
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	dalamL1 = didalam(cx1, cy1, r1, x, y)
	dalamL2 = didalam(cx2, cy2, r2, x, y)

	if dalamL1 && dalamL2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}

}
