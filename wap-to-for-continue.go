/*
*@author: mr-black
*@desc: in for loop init and post statement are optional
**/

package main

import "fmt"

func main(){
	count:=1
	for ;count<=10; {
		fmt.Printf("Count: %v\n", count)
	}
}
