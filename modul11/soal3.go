package main

import "fmt"

func main() {
	var p tabPartai
	var n, x, idx int
	n = 0
	fmt.Scan(&x)
	for x != -1 {
		idx = posisi(p, n, x)
		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&x)

	}
}
