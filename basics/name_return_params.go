package main

import (
	"fmt"
)
// name return parameterised function
func namedret() (name string, rollno int) {
	name= "Anuj"
	rollno=10
	return
}

func main()  {
	nma, rna := namedret()
	fmt.Printf("name: %v password: %v", nma, rna)
}