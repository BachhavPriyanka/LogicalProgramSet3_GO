package main
import "fmt"

func main(){
	var num int
	fmt.Println("Enter number :")
	fmt.Scan(&num)
	multiples(num)
}

func multiples(n int){
	sum := 0
	for i := n-1; i >= 1; i-- {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
			fmt.Println(i, "")
		}
	}
}