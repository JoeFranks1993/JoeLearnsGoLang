package main

import "fmt"

type customStruct struct {
	member1 string
	member2 string
	member3 string
}

func main() {
	var x = "This is a string"
	var y = 500
	var z = 5.43234
	var myType = customStruct{member1: "one", member2: "two", member3: "three"}

	var xType = fmt.Sprintf("%T", x)
	var yType = fmt.Sprintf("%T", y)
	var zType = fmt.Sprintf("%T", z)
	var customType = fmt.Sprintf("%T", myType)

	fmt.Println(xType)
	fmt.Println(yType)
	fmt.Println(zType)
	fmt.Println(customType)
}
