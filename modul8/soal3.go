package main
import "fmt"

func main(){
	var x,y int
	fmt.Scan(&x,&y)
	fmt.Println(LCM(x,y))
}

func RecursiveGCD(a,b int) int {
	if b == 0 {
		return a
	}
	return RecursiveGCD(b, a%b)
}

func LCM(a,b int) int {
	// rumus kpk => lcm(a,b) = (a*b)/gcd(a,b)
	if (a == 0 || b == 0) {
		return 0
	}
	return (a*b)/RecursiveGCD(a,b)
}