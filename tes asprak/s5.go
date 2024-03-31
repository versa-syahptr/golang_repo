package main

import "fmt"

type ArrInt [100]int

func main(){
	var x,n,m int
	var arr ArrInt
	fmt.Scan(&x)
	for i:=0; i<x; i++ {
		fmt.Scan(&m, &n)
		for i:=0; i<n; i++ {
			fmt.Scan(&arr[i])
		}
		eskrim(arr, n, m)
	}
}


func eskrim(arr ArrInt, n,m int){
	var a,b, i,j int
	for i=0; i < n && a+b != m ; i++ {
		a = arr[i]
		for j=n-1; j > i; j-- {
			b = arr[j]
			if a+b == m {
				fmt.Println(i, " ", j)
				break
			}
		}
	}
	
}