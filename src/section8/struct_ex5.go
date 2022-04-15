package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calcuate() float64 {
	return e.salary + e.bonus
}

//같은 메소드이지만 Executives 구조체로 재정의
func (e Executives) Calcuate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee     // is a 관계 (=임원도 직원이다)
	specialBonus float64
}

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴

	//예제1
	//직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	//임원
	ex := Executives{
		Employee{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 kim: ", int(ep1.Calcuate()))
	fmt.Println("ex1 park: ", int(ep2.Calcuate()))

	//Employee 구조체 통해서 메소드 호출
	fmt.Println("ex2 lee: ", int(ex.Calcuate()))
	fmt.Println("ex2 lee: ", int(ex.Employee.Calcuate()+ex.specialBonus)) // Employee 접근해서 직접 호출
}
