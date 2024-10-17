package main

import "fmt"

// loop
func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		fmt.Printf("%d * %d = %d \n", result, i, result*(i))
		result *= i

	}
	fmt.Printf("\nFactorial of %d is: %d\n", n, result)
}

func recFactorial(n int) int {
	result := 1
	if n > 0 {
		subResult := recFactorial(n - 1)
		result = n * subResult
		fmt.Printf("%d * %d = %d \n", n, subResult, result)
		return result
	}

	return 1
}

func main() {

	var number int
	fmt.Print("enter number: ")
	fmt.Scan(&number)

	fmt.Println("--------------------------------")
	factorial(number)
	fmt.Println("--------------------------------")
	recFactorial(number)
	fmt.Println("--------------------------------")

}
