package main
import "fmt"

func main(){
	var num = 10
	var num1 = 0
	var num2 = 1
	var count = 0
	var num3 int

	fmt.Print(num1, " ",num2)

	for num1 < num{
		if num1 % 2 == 0{
			count += 1
		}
		num3 = num1 + num2
		fmt.Print(" ",num3)
		num1 = num2
		num2 = num3
	}
	fmt.Println("\nCount of Even Numbers in Fibonacii : ",count)
}

// func c (x int, y int ) int{

// 	if(x/2 == 0){
// 		y +=1
// 	}
// 	return y
// }

// num = 10
// current := 1
// next := 2
// for current < num{

// }

