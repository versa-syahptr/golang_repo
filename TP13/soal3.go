// Buatlah program yang digunakan untuk menghitung digit suatu bilangan secara rekursif!
package main

import "fmt"

func main() {
	number := 123456789
	fmt.Println(digitsCount(number))
}

func digitsCount(x int) int {
	if x > 0 {
		return 1 + digitsCount(x/10)
	} else {
		return 0
	}
}
