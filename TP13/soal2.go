// Buatlah program yang menampilkan isi elemen array secara rekursif!
package main

import "fmt"

const MAX = 5
type tabInt [MAX]int

func main() {
	var array = tabInt{1,4,2,5,7}
	cetakArrayRecursive(array, MAX, 0)

}

func cetakArrayRecursive(arr tabInt, len, start int){
	if start < len && start >= 0{
		fmt.Println(arr[start])
		cetakArrayRecursive(arr, len, start+1)
	}
}