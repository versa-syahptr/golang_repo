package main

import "fmt"

func main() {
	var text string

	for text != "END" {
		fmt.Scan(&text)
		fmt.Print(text, " ")
	}

}
