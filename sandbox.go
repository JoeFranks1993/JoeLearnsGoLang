package main

import "fmt"

func main() {
	// Test comment for github change
	fmt.Println(Joe(3))
	fmt.Println(AddNums(5, 5))
	fmt.Println(AddNums(GetNums()))
	fmt.Println("Watch this video, it seems like it will help you learn go. " + GetThatVideoYouWereWatchingToLearnGoLang())

	var dog = Dog{color: "White And Yellow"}
	Move(dog)
	EatDinner(dog)

	var cat = Cat{color: "White", weight: 100}
	Move(cat)
	EatDinner(cat)
}

func Joe(input int) (result string, i int) {
	i = input
	result = "Hey Joe, How's it going??"
	return result, i // Or you could just say return and this is implied.
}

func AddNums(x int, y int) (sum int) {
	return x + y
}

///////////////////////////////////////////////////////////////

func GetNums() (int, int) {
	return 10, 7
}

func GetThatVideoYouWereWatchingToLearnGoLang() (url string) {
	return "https://www.youtube.com/watch?v=YS4e4q9oBaU0"
}

//////////////////////////////////////////////

type Animal interface {
	Walk()
	Eat()
}

type Dog struct {
	color string
	age   int
	breed string
}

type Cat struct {
	color  string
	age    int
	weight int
}

func (d Dog) Walk() {
	fmt.Println("The Dog Walked Around On Four Legs")
}

func (c Cat) Walk() {
	fmt.Println("Budda Waddles Over To The Food Dish, BIIIIIG OL' GURL")
}

func (d Dog) Eat() {
	fmt.Println("Dog Eats Food")
}

func (c Cat) Eat() {
	fmt.Println("Budda Eats like she has never seen food in her life before because she's BIIIIIG")
}

func Move(a Animal) {
	a.Walk()
}

func EatDinner(a Animal) {
	a.Eat()
}

// That video you were watching
//https://www.youtube.com/watch?v=YS4e4q9oBaU
