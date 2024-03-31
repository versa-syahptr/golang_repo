package main

import "fmt"

func main(){
	var nilai, n, jumlah int
	var rerata float64

	fmt.Scan(&nilai)

	for nilai != -1{
		n++
		jumlah += nilai
		fmt.Scan(&nilai)
	}
	if n != 0 {
		rerata = float64(jumlah) / float64(n)
	}
	fmt.Println("rata-rata:", rerata)
}