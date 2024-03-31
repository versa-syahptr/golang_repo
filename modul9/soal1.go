package main

import "fmt"

type str [2020]rune

func main() {
	var s rune
	var teks str

	len := 0
	fmt.Scanf("%c", &s)
	for s != '.' {
		teks[len] = s
		len++
		fmt.Scanf("%c", &s)
	}

	palindrom := true
	for i:=0; i< len && palindrom; i++ {
		palindrom = teks[i] == teks[len-i-1]
	}
	fmt.Println(palindrom)
}