package main

import (
	"fmt"
)

func main() {
 name:="Anuj"
 rollno:=10
 fmt.Printf("name = %v rollno = %v \n", name, rollno)
 call_by_value(&name, &rollno)
 fmt.Printf("name = %v rollno = %v \n", name, rollno)
}

func call_by_value(name *string, rollno *int)  {
	*name = *name + " Singh"
	*rollno = *rollno+10
}