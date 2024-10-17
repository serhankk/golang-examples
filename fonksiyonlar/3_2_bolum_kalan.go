package main

import "fmt"

func divider(a, b int) (int, int) {
	var result int = a / b
	var remainder int = a % b

	return result, remainder
}

func main() {
	var firstNum int
	var secNum int

	fmt.Scan(&firstNum, &secNum)
	// fmt.Println(firstNum, secNum)
	myResult, myRemainder := divider(firstNum, secNum)

	fmt.Printf("result: %d\n", myResult)
	fmt.Printf("remainder: %d\n", myRemainder)
}
