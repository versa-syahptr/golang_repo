package main
import "fmt"

type ArrInt [100]int

func main(){
	var n int
	var arr ArrInt
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}
	sort(arr, n)
}

func sort(arr ArrInt, n int){
	var i, j int
	for i=1; i<n; i++{
		j = i
		for j > 0 && arr[j-1] > arr[j] {
			swap(&arr[j-1], &arr[j])
			j--
		}
		for x:=0; x<n; x++ {
			fmt.Print(arr[x], " ")
		}
		fmt.Println()
	}
}

func swap(a,b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}