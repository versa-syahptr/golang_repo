package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

func terdekat(T []titik) (float64, titik, titik) {
	jarakmin := 9999.0
	var p1,p2 titik
	n := len(T)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			t1 := T[i]
			t2 := T[j]
			jarak := math.Sqrt(math.Pow(t1.x-t2.x, 2) + math.Pow(t1.y-t2.y, 2))
			if jarak < jarakmin {
				jarakmin = jarak
				p1 = t1
				p2 = t2
			}
		}
	}
	return jarakmin, p1, p2
}

func main() {
	var t1, t2 titik
	var jmin float64
	arr := []titik{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	jmin, t1, t2 = terdekat(arr)
	fmt.Printf("T1(%.f,%.f) T2(%.f,%.f) jarak: %f", t1.x, t1.y, t2.x, t2.y, jmin)
}
