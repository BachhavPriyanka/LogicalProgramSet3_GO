package main
import "fmt"

func main(){
	var firstNum, secondNum int
	fmt.Println("Enter First number : ")
	fmt.Scanln(&firstNum)
	fmt.Println("Enter Second number : ")
	fmt.Scanln(&secondNum)

	result := sum(firstNum, secondNum)
	fmt.Println("The sum of Two numbers :",result)

	result1 := diff(firstNum, secondNum)
	fmt.Println("The difference between Two numbers :",result1)

	result2, res := getSum_n_Diff(firstNum, secondNum)
	fmt.Println(result2, res)
}

func sum(a int, b int) int{
	return a + b
}

func diff(x int, y int) int{
	var difference int
	if x > y{
		difference = x - y
	} else {
		difference = y - x
	}
	return difference
}

func getSum_n_Diff(a int, b int) (int, int){
	var diff int
	sum := a + b
	if a > b{
		diff = a - b
	} else {
		diff = b - a
	}
	return sum, diff
}