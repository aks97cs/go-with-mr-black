/*
*@Author: mr-black
*@Desc: program for 2-d mat.
**/

package main
import "fmt"

func main(){
	var ar [3][3]int
	fmt.Println(ar)

	// matrix input

	for i:=0; i<3; i++{
		for j:=0; j<3; j++{
			fmt.Scan(&ar[i][j])
		}
	}

	fmt.Println("Matrix: %v",ar)
}
