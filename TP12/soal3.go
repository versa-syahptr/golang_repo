package main

const N = 100

func CariSekuensial(tab [N]int, v int) int {
	idx := -1

	for i:=0; i<N && tab[i] != 0 && idx == -1; i++ {
		if tab[i] == v {
			idx = i
		}
	}
	return idx
}