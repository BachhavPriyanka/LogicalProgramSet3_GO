package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter a number to find Factorial : ")
	fmt.Scanln(&n)
	result := factorial(n)
	fmt.Println("Factorial :",result)
}

func factorial(num int) int{
	if num == 0{
		return 1
	} else {
		return num * factorial(num - 1)
	}
}