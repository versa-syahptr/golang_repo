package main

import (
	"fmt"
)

type List []interface{}

func (l *List) append (e interface{}){
	*l = append(*l, e)
	fmt.Println("the l:", *l)
}

func main() {
	alis := List{1,2}
	alis.append(3)
	alis.append(4.1)
	fmt.Print(alis)
}