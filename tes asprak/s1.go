package main
import "fmt"

func main(){
	var arr [100]int
	var n int
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}
	rasio(arr, n)
}

func rasio(arr [100]int, n int){
	var pos,neg,zo int
	var rp, rn, rz float64
	for i:=0; i<n; i++ {
		if arr[i] > 0 {
			pos++
		} else if arr[i] < 0 {
			neg++
		} else {
			zo++
		}
	}
	rp = float64(pos)/float64(n)
	rn = float64(neg)/float64(n)
	rz = float64(zo)/float64(n)
	fmt.Printf("%.6f\n%.6f\n%.6f\n", rp, rn, rz)
}