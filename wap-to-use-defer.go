/*
*@author: mr-black
*@desc: use of defer
**/

package main
import "fmt"

func main(){
	// defer statement defers the execution of a function until the surrounding function returns.

	defer fmt.Println("Mr-Black")

	fmt.Println("Hello")
	fmt.Println("Lord:")
}
