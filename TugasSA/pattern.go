package main

import "fmt"

func pola_string(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	var i int

	for i = 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			return i
		}
	}
	fmt.Println("i", i, n-m)
	return -1
}

func main() {
	s := "saya makan nasi"
	p := "kan"
	fmt.Println(pola_string(s, p))
}
