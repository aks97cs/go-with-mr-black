/*
*@author: mr-black
*@desc: program to make calculator
**/

package main
import "fmt"

func main()  {
	fmt.Println("########## Mr-Black #########")
	fmt.Println("Choices are : ")

	fmt.Println("1. Addition")
	fmt.Println("2. Sustraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Println("Mr-Black want you two number to peform the calculation.., After each entry hit the enter..")
	
	// var decalaration
	var num1, num2 int

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	fmt.Println("Mention your choice of operation from the above menu")
	var ch int
	fmt.Scanln(&ch)
	switch ch {
	case 1:
		fmt.Printf("Addition Result : %v", num1+num2)
	case 2:
		fmt.Printf("Substraction Result: %v", num1-num2)
	case 3:
		fmt.Printf("Multiplication Result: %v", num2*num1)
	case 4:
		fmt.Printf("Division Result: %v", num1/num2)
	default:
		fmt.Println("You have entered a wrong choice")
		
	}
}