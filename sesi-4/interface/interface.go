package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type reactangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r reactangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r reactangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

//V.1
// func main() {
// 	var c1 shape = circle{radius: 5}
// 	var r1 shape = reactangle{width: 3, height: 2}

// 	fmt.Printf("type of c1: %T\n", c1)
// 	fmt.Printf("type of r1: %T\n", r1)

// 	fmt.Println("Circle Area:", c1.area())
// 	fmt.Println("Circle Perimeter:", c1.perimeter())

// 	fmt.Println("Reactangle Area:", r1.area())
// 	fmt.Println("Reactangle Perimeter:", r1.perimeter())
// }

//V.2
func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = reactangle{width: 3, height: 2}

	c1.(circle).volume() //type assertion-> mengembalikan ke type data aslinya (struct)
	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle value: %v\n", value)
		fmt.Printf("Circle volume: %f\n", value.volume())
	}
	print("Reactangle", r1)
	print("Circle", c1)

}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

//TYPE ASSERTION

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
