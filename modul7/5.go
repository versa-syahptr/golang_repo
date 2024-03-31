package main

import "fmt"

func main() {
	var nama, nmax, nmin string
	var i, p int
	var nilai, total float64
	var rerata, reratamin, reratamax, reratalama float64
	for {
		fmt.Scan(&nama)
		if nama == "STOP" {
			break
		}
		fmt.Scan(&p)
		for i = 0; i < 3*p; i++ {
			fmt.Scan(&nilai)
			total += nilai
		}
		rerata = total / float64(3*p)
		if reratalama == 0 {
			reratalama = rerata
		}
		if reratalama < rerata {
			reratamax = rerata
			nmax = nama
		}
		if reratalama > rerata {
			reratamin = rerata
			nmin = nama
		}
		reratalama = rerata
	}
	fmt.Printf("%s %.2f\n", nmax, reratamax)
	fmt.Printf("%s %.2f\n", nmin, reratamin)

}
