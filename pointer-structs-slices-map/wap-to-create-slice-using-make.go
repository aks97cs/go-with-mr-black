/*
*@author: mr-black
*@desc: program to create slice with make
**/
package main
import "fmt"

func main(){
	s:= []int{1,2,3,4}
	fmt.Printf("Original slice: %v\n", s)

	// reducing via slicing ..
	s = s[:2]

	fmt.Printf("Reduced slice: %v\n", s)

	// extending slices
	s = s[:4]
	fmt.Println("Extended slice: %v", s)

	/* ******  now slicing with make ******* 	
	Slices can be created with the built-in make function; 
	this is how you create dynamically-sized arrays. */

	ar := make([]int, 5)
	fmt.Printf("Array with make : %v\n", ar)

	// specify capacity using make
	c_ar := make([]int, 2, 5)
	fmt.Printf("Array with capacity using make: %v %v\n", c_ar, cap(c_ar))
}
