package main
import "fmt"

func main(){
	var jmlSKS, jmlNilai, i, n, sks int
	var indeks string
	fmt.Scan(&n)
	
	for i = 1; i <= n; i++ {
		fmt.Scan(&indeks, &sks)
		for indeks != "A" && indeks != "B" && indeks != "C" && 
			indeks != "D" && indeks != "E" || sks <= 0 {
			fmt.Scan(&indeks, &sks)
		}
		jmlSKS += sks
		switch indeks {
		case "A":
			jmlNilai += 4 * sks
		case "B":
			jmlNilai += 3 * sks
		case "C":
			jmlNilai += 2 * sks
		case "D":
			jmlNilai += 1 * sks
		case "E":
			jmlNilai += 1 * sks
		}
	}
	fmt.Println(float64(jmlNilai)/float64(jmlSKS))
}
