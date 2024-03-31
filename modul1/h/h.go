package main

import "fmt"

func main(){
	var jari2, tinggi int

	fmt.Scan(&jari2, &tinggi)
	fmt.Println(hitungVolume(jari2, tinggi))
}

func hitungVolume(r,t int) float64{
	var pi float64 = 3.14
	return float64(r*r)*pi*float64(t)
}