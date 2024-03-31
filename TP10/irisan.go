package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	/* mengembalikan true apabila bilangan val ada di dalam array T yang berisi
	sejumlah n bilangan bulat */
	var found bool
	for i := 0; i < n && !found; i++ {
		found = T[i] == val
	}
	return found
}

func inputSet(T *set, n *int) {
	/* I.S. data himpunan telah siap pada piranti masukan
	   F.S. array T berisi sejumlah n bilangan bulat yang berasal dari masukan
	   (masukan  berakhir apabila bilangan ada yang duplikat, atau array penuh)
	   Catatan: Panggil function exist di sini untuk membantu pengecekan */
	var entry int
	fmt.Scan(&entry)

	for !exist(*T, *n, entry) {
		T[*n] = entry
		*n++
		fmt.Scan(&entry)
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	/* I.S. terdefinisi himpunan T1 dan T2 yang berisi sejumlah n dan m anggota  himpunan
	 	F.S. himpunan T3 berisi sejumlah h bilangan bulat yang merupakan irisan dari
		 himpunan T1 dan T2
	 	Catatan: Panggil function exist di sini untuk membantu pengecekan */

	*h = 0 // memastikan

	// iterasi pada T1
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) { // cek pada T2
			T3[*h] = T1[i]
			*h++
		}
	}
}

func printSet(T set, n int) {
	/* I.S. terdefinisi sebuah himpunan T yang berisi sejumlah n bilangan bulat
	F.S. menampilkan isi array T secara horizontal (dipisahkan oleh spasi) */
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println() // line break
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int
	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
