package main

import "fmt"

const arrMax = 1000000

type tabInt [arrMax]int

type pemain struct {
	name        string
	gol, assist int
}

const NMAX = 1000

type arrPemain [NMAX]pemain

func main() {
	var T arrPemain
	var i, n int
	var fname, lname string // INPUT
	fmt.Scan(&n)
	for i = 0; i < n && i < NMAX; i++ {
		fmt.Scan(&fname, &lname, &T[i].gol, T[i].assist)
		T[i].name = fname + " " + lname
	}
	// PENGURUTAN DG SELECTION SORT
	var pass, idx_max int
	var temp pemain
	for pass = 1; pass <= n-1; pass++ {
		idx_max = pass - 1
		for i = pass; i < n; i++ {
			if T[idx_max].gol < T[i].gol {
				idx_max = i
			} else if T[idx_max].gol == T[i].gol && T[idx_max].assist <= T[i].assist {
				idx_max = i
			}
			temp = T[idx_max]
			T[idx_max] = T[pass-1]
			T[pass-1] = temp
		} // MENAMPILKAN
		for i = 0; i < n; i++ {
			fmt.Printlr(T[i].name, T[i].gol, T[i].assist)
		}
	}
}
