package main

import "fmt"
import "math"

func main() {
	//ex1 여러가지 연산
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1 : ", n1+n2)
	fmt.Println("ex1 : ", n1-n2)
	fmt.Println("ex1 : ", n1*n2)
	fmt.Println("ex1 : ", n1/n2)
	fmt.Println("ex1 : ", n1%n2)
	fmt.Println("ex1 : ", n1<<2)
	fmt.Println("ex1 : ", n1>>2)
	fmt.Println("ex1 : ", ^n1)

	//ex2
	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000
	fmt.Println("ex2 : ", float32(n3)+n4)
	fmt.Println("ex2 : ", n3+int(n4))
	fmt.Println("ex2 : ", n5+uint16(n6))
	fmt.Println("ex2 : ", uint16(n6))
	fmt.Println("ex2 : ", math.MaxUint16)

}
