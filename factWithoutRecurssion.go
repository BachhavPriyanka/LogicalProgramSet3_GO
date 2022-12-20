package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter a number to find Factorial : ")
	fmt.Scanln(&n)
	factorial(n)
	
}

func factorial(num int){
	var fact int
	fact = 1
	for num >= 1 {
		fact *= num
		num -= 1
	}
	fmt.Println("Factorial :",fact)
}