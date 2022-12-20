package main

import "fmt"

func nonPrimeOdd(num int) {
	count := 0
	for i := num - 1; i > 1; i-- {
		result := isNumPrime(i)
		if result == false {
			if i%2 != 0 {
				fmt.Printf("Non Prime Odd : %d\n", i)
				count++
			}
		}
	}
	fmt.Printf("Count : %d ", count + 1)
	fmt.Println("")
}

func isNumPrime(num int) bool {
	isPrime := 0
	if num == 0 {	
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = 1
				return false
			}
		}
		if isPrime == 0 {
			return true
		}
	}
	return false
}

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)
	nonPrimeOdd(num)
}

