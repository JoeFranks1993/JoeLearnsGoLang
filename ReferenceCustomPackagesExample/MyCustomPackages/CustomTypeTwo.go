package custompackage

import "fmt"

type CustomTypeTwo struct {
	Member1 string
	Member2 string
	Member3 int
}

func (x CustomTypeTwo) CustomMethodTwo() {
	fmt.Println("This Is A Method Invoked By Custom Struct Two.")
}
