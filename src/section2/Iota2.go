package main

import "fmt"

func main() {
	//열거형2 Iota 사용
	const (
		A = iota * 10
		B
		C
	)
	fmt.Println(A, B, C)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)
	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)
}
