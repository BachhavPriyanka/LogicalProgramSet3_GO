package main
import "fmt"

func main(){
	var num = 5
	var num1 = 0
	var num2 = 1
	fmt.Print(num1, " ",num2)
	fib(num1, num2, num-2)
}

func fib(n1 int, n2 int, n int){
	var num3 int
	if n > 0{
		num3 = n1 + n2
		fmt.Print(" ",num3)
		n1 = n2
		n2 = num3
		
		fib(n1,n2,n-1)
	}
}