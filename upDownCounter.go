package main

import "fmt"

func main() {
	//Insert your code here
	// infinite for loop like while (true) {}
	var num1, num2 int
	var dummy string

	fmt.Println("Enter a 2 numbers")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%s", &dummy)
	fmt.Scanf("%d", &num2)
	fmt.Scanf("%s", &dummy)

	// flip the numbers, so num2 > num1
	if num1 > num2 {
		num1, num2 = num2, num1
	}

	fmt.Println("Ascending")
	for x := num1; x <= num2; x++ {
		fmt.Printf("%d\n", x)
	}

	fmt.Println("Descending")
	for x := num2; x >= num1; x-- {
		fmt.Printf("%d\n", x)
	}

}
