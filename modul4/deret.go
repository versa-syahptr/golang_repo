package main

import "fmt"

func main(){
	var n,a,r,i int
	fmt.Scan(&n,&a,&r)
	fmt.Print(0)

	
	for i = 1; i <= n; i++{
		fmt.Print(" + ", a*i*r)
	}
}