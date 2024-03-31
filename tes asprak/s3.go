package main
import "fmt"

func main(){
	var x float64
	fmt.Scan(&x)
	fmt.Printf("%.f",factorial(x))
}

func factorial(n float64) float64{
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

