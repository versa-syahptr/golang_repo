package main

import "fmt"

func main(){
	var n, jumlah, t1,t2,t3 int

	fmt.Scan(&n)

	for i:=0; i < n; i++ {
		fmt.Scanf("%d %d %d", &t1, &t2, &t3)
		if t1 + t2 + t3 >= 0 {
			jumlah++
		}
	}
	fmt.Println(jumlah)
}