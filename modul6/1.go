package main
import "fmt"

func findFactorial(n int, res *int) {
	var j int
	*res = 1
	for j = 1; j <= n; j++{
		*res *= j
	}
}

func combination(n,r int) int {
	var nf, nrf,rf int
	findFactorial(n, &nf)
	findFactorial(n-r, &nrf)
	findFactorial(r, &rf)
	return nf / (rf * nrf)
}

func permutation(n,r int) int {
	var nf,nrf int
	findFactorial(n, &nf)
	findFactorial(n-r, &nrf)
	return nf / nrf
}

func main(){
	var a,b,c,d int
	var p1,c1,p2,c2 int
	fmt.Scan(&a,&b,&c,&d)
	p1 = permutation(a,c)
	c1 = combination(a,c)
	p2 = permutation(b,d)
	c2 = combination(b,d)
	fmt.Println(p1,c1)
	fmt.Println(p2,c2)
	
}