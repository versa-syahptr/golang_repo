package main

const N = 100

func NilaiMinimum(tab [N]int, n int) int {
	var min_idx int
	for i:=1; i<n; i++ {
		if tab[min_idx] > tab[i] {
			min_idx = i
		}
	}
	return min_idx
}