package main

const N = 100

func TerurutB(tab *[N]int, n int) {
	for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if tab[j] < tab[i] {
                tab[j], tab[i] = tab[i], tab[j]
            }
        }
    }
}