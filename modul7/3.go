package main

import "fmt"

func main(){
	var bin, dec int
	fmt.Scan(&bin)
	konversi(bin, &dec)
	fmt.Println(dec)
}

func pangkat(a,b int) int{
	/* Mengembalikan nilai a dipangkatkan b (0 <= a,b <= 10), 
	gunakan operasi perkalian di dalam perulangan, misalnya 2 pangkat 4 artinya 2 x 2 x 2 x 2 */
	var i, res int
	res = 1
	for i=0; i<b; i++{
		res *= a
	}
	return res
}

func konversi(biner int, desimal *int) {
	/* I.S. terdefinisi bilangan bulat yang merepresentasikan bilangan biner
	F.S. desimal berisi hasil konversi dari bilangan biner */
	var ld, exp, res int
	for biner > 0 {
		ld = biner%10  // last digit
		res += ld * pangkat(2,exp)
		biner /= 10 
		exp++  // exponen
	}
	*desimal = res
}