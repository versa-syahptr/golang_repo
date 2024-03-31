package input
import "fmt"

func Input(prompt string) string {
	var i string
	fmt.Print(prompt)
	fmt.Scanln(&i)
	return i
}