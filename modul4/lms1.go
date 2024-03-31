package main

import "fmt"

func main(){
	var i,n,a,b, res int
	fmt.Scan(&n)

	for i=0; i < n; i++{
		fmt.Scan(&a, &b)
		res = 1
		for j := 1; j <= b; j++{
			// fmt.Print(j)
			res = res * a
		}
		fmt.Println(res)
	}
}