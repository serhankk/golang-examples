package main

import "fmt"

func main() {

	var myInt int = 100
	var myDecimal float64 = 12.5
	var myBool bool = true
	var myString string = "Hello, World!"

	//Addition
	addition := myInt + myInt
	fmt.Printf("addition :  %d + %d = %d\n", myInt, myInt, addition)

	//Subtraction
	substraction := myDecimal - myDecimal
	fmt.Printf("subtraction : %.2f - %.2f = %.2f\n", myDecimal, myDecimal, substraction)

	//Multiplication
	multiplication := myDecimal * myDecimal
	fmt.Printf("multiplication : %.2f * %.2f = %.2f\n", myDecimal, myDecimal, multiplication)

	//Division
	division := myInt / myInt
	fmt.Printf("division : %d / %d = %d\n", myInt, myInt, division)

	//Modulus
	modulus := myInt % 2
	fmt.Printf("remainder : %d %% %d = %d\n", myInt, myInt, modulus)

	// ----------------------------------------------------------------
	fmt.Printf("my bool's value: %t \n", myBool)
	fmt.Printf("my string's value: %s \n", myString)

}
