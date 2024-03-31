package main

import "fmt"

func main(){
	var t1,t2,t3,t4 uint  // unsigned integer (integer positif)
	var h1,h2,h3,h4 bool

	fmt.Scan(&t1,&t2,&t3,&t4)

	h1 = t1%400 == 0 || (t1%4 ==0 && t1%100 != 0)
	h2 = t2%400 == 0 || (t2%4 ==0 && t2%100 != 0)
	h3 = t3%400 == 0 || (t3%4 ==0 && t3%100 != 0)
	h4 = t4%400 == 0 || (t4%4 ==0 && t4%100 != 0)

	fmt.Println(h1 && h2 && h3 && h4)
}
