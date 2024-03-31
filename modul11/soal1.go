package main

import "fmt"

const arrMax = 1000000

type tabInt [arrMax]int

func main(){
	var array tabInt
	var x int
	fmt.Scan(&x)
	i:=0
	for i < arrMax && x != -5313541{
		if x == 0 {  // urutkan & cari median
			fmt.Println("I:", i)
			sort(&array, i)
			printArray(array, i)
			fmt.Println("median: ", median(array, i))
		} else {
			array[i] = x  // isi array
			i++
		}
		fmt.Scan(&x)

	}
}

func sort(T *tabInt, n int){
	for i := 0; i < n; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if T[j] < T[minIndex] {
				// swap
				temp := T[j]
                T[j] = T[minIndex]
				T[minIndex] = temp
            }
        }
	}
}

func printArray(T tabInt, n int){
	for i:=0; i<n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func median(T tabInt, n int) int{
	var med, medIdx int
	len := n
	medIdx = len/2
	if len%2 == 0 {
		med = (T[medIdx] + T[medIdx-1])/2
	} else {
		med = T[medIdx]
	}
	return med
}