package main

import "fmt"

type char struct {
	s string
	c rune
}

func Char(c rune) char {
	return char{string(c), c}
}

func (c char) add(o char) string {
	return c.s + o.s
}

func main(){
	var a,b,c char
	a = Char('a')
	b = Char('b')
	c = a+b
}