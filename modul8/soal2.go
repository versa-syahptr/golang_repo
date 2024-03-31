package main
import "fmt"

func main(){
	var i,n, A,B int
	var win string
	fmt.Scan(&n)

	for i = 0; i < n; i++{
		ambilSkor(&A,&B)
	}
	fmt.Println("Petinju A:", A, "dan petinju B:", B)

	if A > B {
		win = "Pemenang adalah petinju A"
	} else if B > A {
		win = "Pemenang adalah petinju B"
	} else {
		win = "Draw"
	}

	fmt.Println(win)
}

func ambilSkor(a,b *int) {
	var a1,a2,a3,b1,b2,b3 int
	fmt.Scan(&a1,&a2,&a3,&b1,&b2,&b3)
	*a += a1+a2+a3
	*b += b1+b2+b3
}