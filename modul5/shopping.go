package main
import "fmt"

func main(){
	var p1,p2,p3,p4,p5 int
	var member string
	var cashback, diskon float64

	fmt.Scan(&member,&p1,&p2,&p3,&p4,&p5)

	cashback = 3.1 * float64(p1+p2+p3)
	diskon = 1.7 * float64(p3+p4+p4)

	if member == "yes"{
		cashback += cashback*15.0/100.0
		diskon += diskon*15.0/100.0
	}

	if p1%2==0 && p2%2==0 && p3%2==0 && p4%2==0 && p5%2==0{
		diskon = 0
	}else if p1%2!=0 && p2%2!=0 && p3%2!=0 && p4%2!=0 && p5%2!=0{
		cashback = 0
	}

	if cashback > 35 {
		cashback = 35
	} else if diskon > 50 {
		diskon = 50
	}

	fmt.Println(cashback, diskon)
}