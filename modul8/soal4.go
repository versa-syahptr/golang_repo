package main
import "fmt"
func tebakan(player rune, nilai int) int{
/* mengembalikan angka tebakan dari player yang diperoleh dari masukan, 
masukan akan dilakukan berulang HINGGA tebakan player benar (sama dengan nilai) 
atau telah  menebak  sebanyak  3x  kesempatan.  Catatan:  perhatikan  string  yang ditampilkan seperti pada contoh*/
	var tebak, i int
	for i = 1; i<=3 || tebak == nilai; i++ {
		fmt.Printf("%c - masukan angka tebakan ke-%d: ", player, i)
		fmt.Scan(&nilai)
	}
	return nilai
}

func tukerPemenang(benar bool, winner, player *rune){
/* I.S. terdefinisi benar yang menyatakan tebakan benar atau salah
F.S. nilai winner dan player saling bertukar apabila benar adalah true */
	var temp rune
	if benar {
		temp = *winner
		*winner = *player
		*player = temp
	}
}

func mulaiRonde(ronde int, winner rune, nilai *int){
	/*  I.S. terdefinisi nomor ronde dan pemenang (winner) dari permainan
	F.S. menampilkan ronde, sedangkan nilai berisi angka rahasia yang diperoleh dari masukan, 
	perhatikan string yang ditampilkan seperti pada contoh*/
	fmt.Println("Ronde", ronde)
	fmt.Printf("%c - masukkan angka rahasia: ", winner)
	fmt.Scan(nilai)
}
func main(){
	// deklarasi variabel
	var winner, player rune
	var ronde, nilai, answer int
	// inisialisasi
	winner = 'A'; player = 'B'; ronde = 1
	// memulai ronde pertama...
	mulaiRonde(ronde, winner, &answer)
	for answer != -101 {
		// pemain menebak...
		nilai = tebakan(player, answer)
		// tuker pemenang apabila syarat memenuhi...
		tukerPemenang(answer==nilai, &winner, &player)
		// tampilakan pemenangnya dan increment ronde
		fmt.Printf("%c adalah pemenangnya\n\n",winner)
		ronde = ronde + 1// mulai ronde selanjutnya...
		mulaiRonde(ronde, winner, &answer)
		}
		fmt.Println("Permainan Selesai")
	}