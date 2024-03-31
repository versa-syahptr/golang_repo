package main
import "fmt"

func main(){
	var b, faktor, nfaktor int
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	faktor = 1
	for faktor <= b{
		if b % faktor == 0 {
			fmt.Print(faktor, " ")
			nfaktor++
		}
		faktor++
	}
	fmt.Println("\nPrima:", nfaktor == 2)
}