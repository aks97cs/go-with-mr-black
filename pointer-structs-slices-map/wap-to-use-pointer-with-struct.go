/*
*@author: mr-black
*@desc: program to use struct with pointer
**/

package main
import "fmt"

type NumberSystem struct {
	count int
}

func main()  {
	number_node := NumberSystem{}
	number_node_ptr := & number_node
	// (*number_node_ptr).count = 100   this is also true
	number_node_ptr.count =100
	fmt.Printf("%d", number_node)
}