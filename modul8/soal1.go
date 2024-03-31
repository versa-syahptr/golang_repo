package main
import "fmt"

func main(){
	var d1, d2, kemunculan int
	fmt.Scan(&d1, &d2)
	for !(genap(d1) && genap(d2)) {  	// demorgan !(a&&b) == !a || !b
		if !(genap(d1) || genap(d2)) { 	// demorgan !(a||b) == !a && !b
			kemunculan++
		}
		fmt.Scan(&d1, &d2)
	}
	fmt.Println(kemunculan)
}

func genap(x int) bool {
	return x % 2 == 0
}