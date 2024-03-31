package main

type Trec struct {
	v1 int
	vx tvx
	v4 int
}

type tvx struct {
	v2, v3 int
}

func BanyakNilai(rec *Trec) {
	var jumlah, min_idx, max_idx int
	tarr := [4]int{rec.v1, rec.vx.v2, rec.vx.v3, rec.v4} // array helper

	for i := 0; i < 4; i++ {
		jumlah += tarr[i]

		if tarr[min_idx] > tarr[i] { // update min index
			min_idx = i
		}
		if tarr[max_idx] < tarr[i] { // update max index
			max_idx = i
		}
	}
	rec.v1 = tarr[min_idx]
	rec.vx.v2 = jumlah
	rec.vx.v3 = jumlah / 4 // integer div
	rec.v4 = tarr[max_idx]
}
