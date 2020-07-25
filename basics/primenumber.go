package main
import (
	"fmt"
)
func main() {
	var num int
	var num_ptr int
	check_flag:=0  // declaring variable 

	fmt.Println("Enter the number:")
	fmt.Scanln(&num)
	
	for num_ptr < num {

		 if num_ptr == 0 || num_ptr == 1{
			 num_ptr++
			 continue
		 }
		
		 if num%num_ptr==0 {
			 check_flag = 1
		 }
		 num_ptr++
	}
	if check_flag ==0{
		fmt.Println("prime number")
	}else{

		fmt.Println("not a prime number")
	}
}