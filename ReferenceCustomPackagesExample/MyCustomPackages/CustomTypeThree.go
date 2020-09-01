package custompackage

import "fmt"

type CustomTypeThree struct {
	Member1 string
	Member2 string
	Member3 int
}

func (x CustomTypeThree) CustomMethodThree() {
	fmt.Println("This Is A Method Invoked By Custom Struct Three.")
}
