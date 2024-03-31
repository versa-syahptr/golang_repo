package main

const N = 100

func CariCepat(tab [N]int, n,v int) int {
	var l,m,r, found int
	l = 0
	r = n-1
	found = -1
	for l <= r && found == -1 {
		m = (l+r)/2
		if tab[m] > v {
			l = m+1
		} else if tab[m] < v{
			r = m-1
		} else {
			found = m
		}
	}
	return found
}