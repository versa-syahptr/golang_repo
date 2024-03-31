package main

import "fmt"

func main(){
	var g1, g2, g3, g4 float64
	var kg1, pon1, ons1 float64
	var kg2, pon2, ons2 float64
	var kg3, pon3, ons3 float64
	var kg4, pon4, ons4 float64

	fmt.Scan(&g1,&g2,&g3,&g4)

	kg1 = g1/1000.0; pon1 = g1/453.592; ons1 = g1/28.3495
	kg2 = g2/1000.0; pon2 = g2/453.592; ons2 = g2/28.3495
	kg3 = g3/1000.0; pon3 = g3/453.592; ons3 = g3/28.3495
	kg4 = g4/1000.0; pon4 = g4/453.592; ons4 = g4/28.3495

	fmt.Println(kg1, pon1, ons1)
	fmt.Println(kg2, pon2, ons2)
	fmt.Println(kg3, pon3, ons3)
	fmt.Println(kg4, pon4, ons4)


}