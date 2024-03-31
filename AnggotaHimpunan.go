package main

import "fmt"

var i int

func main() {
	i = 0
}

func isValid(T himpunan, n int) bool {
	var i, j int
	var valid bool
	valid = true
	i = 0
	j = 0
	for i < n && valid {
		j = i + 1
		for j < n {
			if T[i] == T[j] {
				valid = false
			}
			j = j + 1
		}
		i = i + 1
	}
	return valid
}
