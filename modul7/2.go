package main

import "fmt"

func main(){
    // deklarasi variabel
	var hari1, bulan1, hari2 string
	var tanggal1, tahun1, tanggal2, tahun2, bln2int int
    // pembacaan masukan
    fmt.Scan(&hari1, &tanggal1, &bulan1, &tahun1) 
    // panggil subprogram untuk penentuan tanggal pengambilan
	pengambilan(tanggal1, angkaBulan(bulan1), tahun1, hari1, &tanggal2, &bln2int, &tahun2, &hari2)
    // tampilkan tanggal pengambilan visa
	fmt.Println(hari2, tanggal2, bulanAngka(bln2int), tahun2)
}

func kabisat(tahun int) bool {
// Mengembalikan true apabila tahun adalah kabisat, false apabila sebaliknya.
    return (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0
}

func angkaBulan(bulan string) int {
/* Mengembalikan angka berdasarkan urutan nama bulan pada kalender masehi (1 hingga 12).
   0 untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua. Contoh: "januari" menjadi 1 */
    switch bulan {
	case "januari":
		return 1
	case "februari":
		return 2
	case "maret":
		return 3
	case "april":
		return 4
	case "mei":
		return 5
	case "juni":
		return 6
	case "juli":
		return 7
	case "agustus":
		return 8
	case "september":
		return 9
	case "oktober":
		return 10
	case "november":
		return 11
	case "desember":
		return 12
	default:
		return 0
	}
}

func bulanAngka(angka int) string {
/* Mengembalikan nama bulan berdasarkan urutan angka bulan pada kalender masehi (1 hingga 12).
   "invalid" untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua. Contoh: 1 menjadi "januari" */
	switch angka {
	case 1:
		return "januari"
	case 2:
		return "februari"
	case 3:
		return "maret"
	case 4:
		return "april"
	case 5:
		return "mei"
	case 6:
		return "juni"
	case 7:
		return "juli"
	case 8:
		return "agustus"
	case 9:
		return "september"
	case 10:
		return "oktober"
	case 11:
		return "november"
	case 12:
		return "desember"
	default:
		return "invalid"
   }
}

func jumlahHari(bulan, tahun int) int {
/* Mengembalikan jumlah hari berdasarkan bulan dan tahun yang terdefinisi,
hati-hati pada bulan februari pada saat kabisat. -1 apabila bulan tidak valid */
    if bulan == 2 && kabisat(tahun) {
		return 29
	} else {
		switch bulan {
		case 2:
			return 28
		case 1,3,5,7,8,10,12:
			return 31
		case 4,6,9,11:
			return 30
		default:
			return -1
		}
	}
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int){
/* I.S. terdefinisi hari1 yang menyatakan hari pengajuan string, asumsi huruf kecil semua
   F.S. hari2 berisi hari pengambilan dan durasi berisi lama pembuatan visa, karena sabtu dan minggu tidak dihitung */
	*durasi = 2
    switch hari1 {
	case "senin":
		*hari2 = "rabu"
	case "selasa":
		*hari2 = "kamis"
	case "rabu":
		*hari2 = "jumat"
	case "kamis":
		*hari2 = "senin"
		*durasi += 2
	case "jumat":
		*hari2 = "selasa"
		*durasi += 2
	}
}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
/* I.S. terdefinisi waktu pengajuan visa, yaitu tanggal1, bulan1, tahun1 dan hari1
   F.S. tanggal2, bulan2, tahun2 dan hari2 berisi waktu pengambilan visa */
    // tentukan durasi, hari pengambilan, dan jumlah hari pada bulan pengajuan
	var durasi, enddate, bln2, jml_hari_bln_ini int
	jml_hari_bln_ini = jumlahHari(bulan1, tahun1)
	mencariDurasi(hari1, hari2, &durasi)  // hari2 disini adalah pointer, jadi tidak perlu ditulis &hari2
	
	
    // dapatkan tanggal pengambilan, asumsi bulan pengambilan dan tahun pengambilan sama dengan waktu pengajuan
	enddate = tanggal1 + durasi
	bln2 = bulan1
	*tahun2 = tahun1
    // cek apabila tanggal pengambilan melebihi lama hari, update tanggal dan bulan pengambilan dengan yang seharusnya
	if enddate > jml_hari_bln_ini {
		enddate %= jml_hari_bln_ini
		bln2++ // +1
	}

	if bln2 > 12 {
		bln2 = 1
		*tahun2++
	}

	*tanggal2 = enddate
	*bulan2 = bln2
    
}

