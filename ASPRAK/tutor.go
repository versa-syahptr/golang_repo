package main

import "fmt"

// deklarasi tipe NumPlay di sini
type NumPlay [100]int

func main() {
	// deklarasi variabel yang dibutuhkan di sini
	var tabA, tabB, arrARev, arrBRev, ganjil1, genap1, ganjil2, genap2 NumPlay
	var n int
	// 1. buatlah dua buah array NumPlay disini
	// isilah array pertama dengan bilangan 3 6 9 12 9 6 3
	isiArray(&tabA, &n)
	// isilah array kedua dengan bilangan 28 21 14 7 14 21 29
	isiArray(&tabB, &n)
	// 2. Tampilkan apakah array tersebut saling identik atau tidak

	// array pertama dengan array pertama
	fmt.Println(identik(tabA, tabA, n))
	// array kedua dengan array kedua
	fmt.Println(identik(tabB, tabB, n))
	// array pertama dengan array kedua
	fmt.Println(identik(tabA, tabB,n))
	// 3. Kembalikan isi array pertama dalam posisi terbalik dan masukan kedalam variabel arrARev
	reverseArray(tabA, &arrARev, n)
	// 4. Kembalikan isi array kedua dalam posisi terbalik dan masukan kedalam variabel arrBRev
	reverseArray(tabB, &arrBRev, n)
	// 5. Tampilkan apakah kedua array tersebut apakah palindrom atau tidak
	fmt.Println(palindrom(tabA, n))
	fmt.Println(palindrom(tabB, n))
	// 6. Pecah array pertama menjadi array ganjil1 dan genap1
	splitArray(tabA, &ganjil1, &genap1, n)
	// 7. Pecah array kedua menjadi array ganjil2 dan genap2
	splitArray(tabB, &ganjil2, &genap2, n)
	// 8. Tampilkan isi array arrA, arrB, arrARev, arrRevB, ganjil1, ganjil2, genap1 dan genap2
	cetakArray(tabA, n)
	cetakArray(tabB, n)
	cetakArray(arrARev, n)
	cetakArray(arrBRev, n)
	cetakArray(ganjil1, n)
	cetakArray(genap1, n)
	cetakArray(ganjil2, n)
	cetakArray(genap2, n)
}

func isiArray(tabA *NumPlay, n *int) {
	/* I.S terdefinisi array tabA yang masih kosong
	   Proses: menerima nilai n dari masukan, kemudian array tabA diisi berdasarkan masukan
	   dari keyboard
	   F.S array tabA terisi sejumlah n bilangan bulat */
	var i, x int
	fmt.Scan(n)
	for i= 0; i < *n; i++{
		fmt.Scan(&x)
		tabA[i] = x
	}
}

func identik(tabA, tabB NumPlay, n int) bool {
	/* mengembalikan true jika setiap elemen tabA identik dengan elemen tabB */
	var i int
	var sama bool = true
	for i = 0; i < n && sama; i++{
		sama = tabA[i] == tabB[i]
	}
	return sama
}

func reverseArray(tabA NumPlay, tabRev *NumPlay, n int) {
	/*I.S terdefinisi array tabA yang sudah berisi elemen
	  F.S array tabRev berisi elemen dari tabA namun dengan posisi terbalik */
	var i,j int
	j = 0
	for i = n-1; i >= 0; i--{
		tabRev[j] = tabA[i]
		j++
	}
}

func palindrom(tabA NumPlay, n int) bool {
	/*mengembalikan true jika setiap elemen tabA identik dengan elemen tabB namun dengan
	  posisi terbalik */
	  var revA NumPlay
	  reverseArray(tabA, &revA, n)
	  return identik(tabA, revA, n)

}

func splitArray(tabA NumPlay, tabGanjil, tabGenap *NumPlay, n int) {
	/* I.S terdefinisi array tabA yang sudah berisi elemen
	   F.S array tabGanjil berisi bilangan ganjil dan array tabGenap berisi bilangan genap
	   yang diperoleh dari array tabA */
	   var i, genap, ganjil int
	   ganjil, genap = 0,0;
	   for i=0; i<n; i++ {
			if tabA[i] % 2 == 0 {
				tabGenap[genap] = tabA[i]
				genap++
			} else {
				tabGanjil[ganjil] = tabA[i]
				ganjil++
			}
	   }

}

func cetakArray(tabA NumPlay, n int) {
	/*I.S terdefinisi array tabA yang sudah terisi sejumlah n bilangan
	  F.S menampilkan isi array ke layar komputer secara horizontal*/
	  var i int
	  for i = 0; i<n; i++{
		fmt.Print(tabA[i], " ")
	  }
	  fmt.Println()

}

