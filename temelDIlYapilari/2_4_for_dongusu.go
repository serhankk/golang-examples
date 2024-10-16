package main

import "fmt"

func main() {

	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var oddNumbers []int
	var evenNumbers []int
	// range function
	for i := range numbers {
		if numbers[i]%2 == 0 {
			evenNumbers = append(evenNumbers, numbers[i])
		} else {
			oddNumbers = append(oddNumbers, numbers[i])
		}
	}
	fmt.Printf("odd numbers: %v\n", oddNumbers)
	fmt.Printf("even numbers: %v\n", evenNumbers)
	fmt.Printf("----------------------------------------------------------------\n")

	oddNumbers2 := []int{}
	evenNumbers2 := []int{}

	// standard loop
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evenNumbers2 = append(evenNumbers2, numbers[i])
		} else {
			oddNumbers2 = append(oddNumbers2, numbers[i])
		}
	}
	fmt.Printf("odd numbers: %v\n", oddNumbers2)
	fmt.Printf("even numbers: %v\n", evenNumbers2)
}
