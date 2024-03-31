package main

import "fmt"

func main(){
	var s1, s2, s3 string

	fmt.Scan(&s1, &s2, &s3)

	fmt.Print(s1==s2 || s1==s3 || s2==s3)
}