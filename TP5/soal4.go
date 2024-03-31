package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var seed int64
	var a,b,c int 
	/* a tebakan user, 
	*  b angka random (yang ditebak pak Pras), 
	*  c angka random juri*/
	fmt.Scan(&seed, &a)
	rand.Seed(seed)
	b = rand.Intn(4)+1 // tebakan Pak Pras
	c = rand.Intn(4)+1 // tebakan juri

	fmt.Println(b)
	fmt.Println(c)

	if a == c && b == c {
		fmt.Println("Seri")
	} else if a == c {
		fmt.Println("Pemenang adalah anda")
	} else if b == c {
		fmt.Println("Pemenang adalah Pak Pras")
	} else {
		fmt.Println("Tidak ada pemenang")
	}

}