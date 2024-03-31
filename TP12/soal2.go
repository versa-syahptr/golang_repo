package main

import "fmt"

const N = 100

func TambahData(tab *[N]int, n *int) {
	var elem int
	fmt.Scan(&elem)
	for *n = 0; *n<N && elem != 9999; *n++ {
		tab[*n] = elem
		fmt.Scan(&elem)
	}
}