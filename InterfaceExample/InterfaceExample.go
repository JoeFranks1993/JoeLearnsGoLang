package main

import "fmt"

func main() {
	var joe = Human {FirstName: "Joseph", LastName: "Franks"};
	var sparky = Dog {Name: "Sparky", Color: "White"};

	Move(joe);
	Eat(joe);

	Move(sparky);
	Eat(sparky);
}


// Interfaces
type Moveable interface {
	Move()
}

type CanEat interface {
	Eat()
}

type ToString interface {
    ToString() string
}

// Methods accepting interface
func Move(a Moveable) {
    a.Move();
}

func Eat(a CanEat) {
    a.Eat();
}

// Types 
type Human struct {
	FirstName string
	LastName string 
	Age int 
	Address string
}

type Dog struct {
	Name string
	Color string
	Age int
	Breed string
}

// Method implementations 
func(h Human) Move() {
    fmt.Println(h.ToString() + " Gets up and walks around the room on two legs")
}

func(h Human) Eat() {
    fmt.Println(h.ToString() + " Has started eating, to do so he uses a fork")
}

func(h Human) ToString() string {
    return h.FirstName + " " + h.LastName;
}

func(d Dog) Move() {
    fmt.Println(d.ToString() + " dog gets up and walks around the yard on four legs")
}

func(d Dog) Eat() {
	fmt.Println(d.ToString() + " eats out of his bowl in the hallway")
}

func(d Dog) ToString() string {
	return d.Name + " The " + d.Color + " Dog";
}