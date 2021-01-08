/*
*@author: mr-black
*@desc: program to take student record
**/

package main
import (
	"fmt"
)

type Student struct{
	name string
	roll_no int
	course_name string
	address string
}

func main(){
	st:= Student{}

	fmt.Printf("Enter the student name: ")
	fmt.Scanf("%s",&st.name)
	
	fmt.Printf("Enter the student roll_no: ")
	fmt.Scanf("%d", &st.roll_no)
	
	fmt.Printf("Enter the student course name: ")
	fmt.Scanf("%s", &st.course_name)
	
	fmt.Printf("Enter student address: ")
	fmt.Scanf("%s", &st.address)

	fmt.Printf("Student record is: %v", st)
}

