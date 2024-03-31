package main

import "fmt"

func main() {
	var text string

	fmt.Scan(&text)

	for text != "selesai" {
		fmt.Println(text)
		fmt.Scan(&text)
	}

}
