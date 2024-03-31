package main

import "fmt"

func main(){
	var penurunan, hari, tdvisit, oldvisit, totalvisit int
	hari = 0
	oldvisit = 0
	for penurunan < 3 {
		hari++
		getInput(hari, &tdvisit)
		for tdvisit > 100 || tdvisit < 0 {
			getInput(hari, &tdvisit)
		}

		if oldvisit > tdvisit {
			penurunan++
		} else {
			penurunan = 0  // reset
		}

		oldvisit = tdvisit
		totalvisit += tdvisit

	}

	fmt.Printf("Museum buka selama %d hari\n", hari)
	fmt.Printf("Rata-rata pengunjung %.2f", float64(totalvisit)/float64(hari))
}

func getInput(iter int, input *int) {
	fmt.Printf("Hari %d : ", iter)
	fmt.Scan(input)
}