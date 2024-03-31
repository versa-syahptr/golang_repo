package main

const N = 100

func TerurutA(T *[N]int, n int){
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