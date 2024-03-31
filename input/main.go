package main

import (
	"fmt"
	"input/input"
)

func main(){
	var text string = input.Input("yes: ")
	fmt.Print(text)
}
