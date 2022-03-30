package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert your code here
	// infinite for loop like while (true) {}
	var num1, num2 string

	fmt.Println("Enter a 2 numbers")
	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	numInt1, _ := strconv.ParseInt(num1, 10, 0)
	numInt2, _ := strconv.ParseInt(num2, 10, 0)

	// flip the numbers, so num2 > num1
	if numInt1 > numInt2 {
		numInt1, numInt2 = numInt2, numInt1
	}

	fmt.Println("Ascending")
	for x := numInt1; x <= numInt2; x++ {
		fmt.Printf("%d\n", x)
	}

	fmt.Println("Descending")
	for x := numInt2; x >= numInt1; x-- {
		fmt.Printf("%d\n", x)
	}

}
