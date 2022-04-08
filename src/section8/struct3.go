package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	color string
	name  string
}

func main() {
	// ex1
	c1 := Car{"red", "220d"}
	c2 := new(Car)
	c2.color, c2.name = "white", "sonata"
	c3 := &Car{}
	c4 := &Car{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(reflect.TypeOf(c1))
	fmt.Println(c2)
	fmt.Println(reflect.TypeOf(c2))
	fmt.Println(c3)
	fmt.Println(reflect.TypeOf(c3))
	fmt.Println(c4)
	fmt.Println(reflect.TypeOf(c4))
	fmt.Println()
}
