package main

import (
	"fmt"
)

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i) //defer func 누적됨 (lifo 후입선출)
	}
}

func main() {
	//ex1
	stack()
}
