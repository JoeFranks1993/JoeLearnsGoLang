package main

import "fmt"

type Cat struct {
	Age int
}

func main() {
	var bebe = Cat{Age: 8}

	fmt.Println("Age Will Not Increment Because The Member Is Not Being Passed As A Reference")
	fmt.Println(bebe.Age) // Output: 8
	incrementAgeIncorrectly(bebe.Age)
	fmt.Println(bebe.Age) // Output: 8
	incrementAgeIncorrectly(bebe.Age)
	fmt.Println(bebe.Age) // Output: 8

	fmt.Println("")

	fmt.Println("Age Will Increment Correctly Because The Member Is Being Passed As A Reference Via A Pointer")
	fmt.Println(bebe.Age) // Output: 8
	incrementAgeCorrectly(&bebe.Age)
	fmt.Println(bebe.Age) // Output: 9
	incrementAgeCorrectly(&bebe.Age)
	fmt.Println(bebe.Age) // Output: 10

}

func incrementAgeIncorrectly(age int) {
	age++
}

func incrementAgeCorrectly(age *int) {
	*age++
}
