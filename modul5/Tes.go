package main
import "fmt"

func main(){
	var a,b,c,d,max,min int
	fmt.Scan(&a,&b,&c,&d)

	switch {
	case a > b && a > c && a > d: max = a
	case b > a && b > c && b > d:
		max = b
	case c > a && c > b && c > d:
		max = c
	case d > a && d > b && d > c:
		max = d
	}
	switch {
	case a < b && a < c && a < d: 
		min = a
	case b < a && b < c && b < d:
		min = b
	case c < a && c < b && c < d:
		min = c
	case d < a && d < b && d < c:
		min = d
	}
	fmt.Println(max, min)
}