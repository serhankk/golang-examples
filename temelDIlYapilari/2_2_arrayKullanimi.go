package main

import "fmt"

func main() {

	// Initialize array
	var myArr [10]int

	fmt.Printf("my array:\t %v\n", myArr)

	// Assign values to array
	myArr[0], myArr[1],
		myArr[2], myArr[3],
		myArr[4] = 10, 20, 30, 40, 50

	// Append new value to array
	myArr[5] = 60
	fmt.Printf("my array:\t %v\n", myArr)
}
