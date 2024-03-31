package main

const N = 100

func NilaiRerata(tab [N]int, n int) int {
	var jumlah int
	for i:=1; i<n; i++ {
		jumlah += tab[i]
	}
	return jumlah/n	//int div
}