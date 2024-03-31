package main

import "fmt"

// a
type nasabah struct {
	kode_asdos, nama, bank string
	rekening int64
}

//b
const maxNasabah = 5
type tabNasabah [maxNasabah]nasabah

// c
func addNasabah(T *tabNasabah, N *int) {
	/*I.S. terdefinisi array T yang berisi data sejumlah N nasabah
	F.S. array T bertambah 1 orang nasabah baru, tampilkan "data penuh" apabila
	array telah penuh*/
	var new nasabah

	if *N < maxNasabah {
		fmt.Scan(&new.kode_asdos, &new.nama, &new.bank, &new.rekening)
		T[*N] = new
		*N++
	} else {
		fmt.Println("data penuh")
	}

}
func cetak(T tabNasabah, N int, X string){
	/*I.S. terdefinisi array T yang berisi data sejumlah N nasabah, dan sebuah
	string X
	F.S. menampilan data nasabah yang menabung pada bank X*/
	var i int
	var cur nasabah
	for i=0; i<N; i++ {
		cur = T[i]
		if cur.bank == X {
			fmt.Print("Kode :", cur.kode_asdos, " , ")
			fmt.Print("Nasabah :", cur.nama, " , ")
			fmt.Print("Bank :", cur.bank, " , ")
			fmt.Println("Rek :", cur.rekening)
		}
	}
}

//d
func main(){
	var i, isi int
	var arrNasabah tabNasabah
	var bank string
	//e
	for i=0; i<10; i++ {
		addNasabah(&arrNasabah, &isi)
	}
	fmt.Scan(&bank)
	cetak(arrNasabah, isi, bank)
}

