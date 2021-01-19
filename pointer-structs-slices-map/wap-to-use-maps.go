/*
*@author: Mr-Black
*@desc: wap to use map in golang
**/

package main
import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main(){

	m = make(map[string]Vertex)
	m["name"] = Vertex{
		40.54, -70.399,
	}

	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["name"].Lat)
}
