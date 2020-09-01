package custompackage

import "fmt"

type CustomTypeOne struct {
	Member1 string
	Member2 string
	Member3 int
}

func (x CustomTypeOne) CustomMethodOne() {
	fmt.Println("This Is A Method Invoked By Custom Struct One.")
}
