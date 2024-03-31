package main

import "fmt"

func main(){
	var status, topi, atk, tali, pisau bool
	var n int
	fmt.Scan(&n)
	status = true
	for i := 0; i < n && status; i++{
		fmt.Scan(&topi, &atk, &tali, &pisau)
		status = topi && atk && tali && pisau
	}
	print(status)
}