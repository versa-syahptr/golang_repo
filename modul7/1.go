package main

import "fmt"

func main(){
	var i,n,a,b, pos,neg int
	fmt.Scan(&n)

	for i=0; i<n; i++ {
		fmt.Scan(&a, &b)

		if a < 0 && b < 0 {
			neg += a + b
		} else if a > 0 && b < 0 {
			pos += a
			neg += b
		} else if a < 0 && b > 0 {
			pos += b
			neg += a
		} else {
			pos += a + b
		}
	}

	fmt.Println("Negative:", neg)
	fmt.Println("Positive:",pos)

}