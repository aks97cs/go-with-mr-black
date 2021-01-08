/*
*@author: mr-black
*@desc: clean way to write long if-then-else chains.
**/

package main
import  (
	"fmt"
	"time"
)

func main(){
	t:=time.Now()
	switch {
	case t.Hour()<12:
		fmt.Println("Good morning Mr-Black")
	case t.Hour()< 17:
		fmt.Println("Good afternoon Mr-Black")
	default:
		fmt.Println("Good evening Mr-Black")
	}
}

