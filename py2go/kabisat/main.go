package main

import "fmt"

func main(){
	var tahun int
	fmt.Scan(&tahun)
	var kabisat bool = tahun % 40 == 0 || tahun % 4 == 0 && tahun % 100 != 0
	fmt.Print(kabisat)
}