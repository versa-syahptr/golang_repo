package main

import "fmt"

func main(){
	var a,b,c,d,e,f, x int

	fmt.Scanf("%d %d %d\n%d %d %d", &a,&b,&c,&d,&e,&f)

	fmt.Scan(&x)

	if x == 0 || x == 360 {
		fmt.Println(a,b,c)
		fmt.Println(d,e,f)
	} else if x == 90 {
		fmt.Println(d,a)
		fmt.Println(e,b)
		fmt.Println(f,c)
	} else if x == 180 {
		fmt.Println(f,e,d)
		fmt.Println(c,b,a)
	} else if x == 270 {
		fmt.Println(c,f)
		fmt.Println(b,e)
		fmt.Println(a,d)
	}
}