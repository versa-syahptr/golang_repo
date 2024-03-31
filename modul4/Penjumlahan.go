package main

import "fmt"

func main(){
	var n, i, bil, sum int
	fmt.Scan(&n)

	for i = 0; i < n; i++{
		fmt.Scan(&bil)
		for bil < 0 || bil > 9 {
			fmt.Scan(&bil)
		}
		sum += bil
	}
	fmt.Print(sum)

}