/*
*@author: mr-black
*@desc: wap-to-append slices
**/

package main
import "fmt"

func main(){
	var sl []int
	fmt.Println(sl)

	// appending slice
	sl = append(sl, 1)
	fmt.Println(sl)

	sl = append(sl, 2)
	fmt.Println(sl)

	sl = append(sl, 3, 4 ,5)
	fmt.Println(sl)

}