package main
import (
	"fmt"
)

func name_and_rollno() (string, int)  {
	return "Anuj", 1
}

func main() {
	name, roll :=name_and_rollno()
	fmt.Printf("name = %v rollno = %v", name, roll)
}