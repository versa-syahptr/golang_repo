package main

import "fmt"

func main(){
	var dec, rem int
	var bin string

	fmt.Scan(&dec)

	for dec > 0 {
		rem = dec%2
		dec = dec/2

		// if rem == 1 {
		// 	bin = "1" + bin
		// } else {
		// 	bin = "0" + bin
		// }
		bin = string(byte(48+rem)) + bin
	}
	fmt.Println(bin) 
}
