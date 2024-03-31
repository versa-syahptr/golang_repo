package main

import (
	"fmt"
)

func main(){
	var gram, kg,gr,hargaKG, hargaGR, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)
	kg = gram/1000
	gr = gram%1000
	fmt.Println("Detail berat:", kg, "kg +", gr, "gr")

	hargaKG = kg*10000
	if gr >= 500{
		hargaGR = gr * 5
	} else {
		hargaGR = gr * 15
	}
	fmt.Println("Detail biaya: Rp.", hargaKG, "+ Rp.", hargaGR)

	if kg > 10 {
		total = hargaKG
	} else {
		total = hargaKG + hargaGR
	}
	fmt.Println("Total biaya: Rp.", total)
}