package main
import (
	"fmt"
)
func main() {
	var num int
	fmt.Println("Enter you number:")
	fmt.Scanln(&num)
	
	if num%2==0{
		fmt.Println("even number")
	}else{
		fmt.Println("odd number")
	}
}