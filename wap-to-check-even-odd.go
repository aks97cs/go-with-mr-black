/*
*@author: mr-black
*@desc: program to check whether a number is even or odd
**/

package main
import "fmt"

func main(){
	fmt.Println("mr-black want's a number from you ..")
	var num int
	fmt.Scanln(&num)

	if num%2==0{
		fmt.Printf("Number: %v is even", num)
	}else{
		fmt.Printf("Number: %v is odd", num)
	}
}
