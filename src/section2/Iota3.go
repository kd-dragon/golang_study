package main

import "fmt"

func main() {
	//열거형 3

	const (
		_ = iota
		A
		_
		C
		D
	)
	fmt.Println("A C D : ", A, C, D)

	// example
	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		_
		PLATINUM
	)

	fmt.Println("DEFAULT : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)
}
