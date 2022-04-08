package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

//일반 메서드
func Price(c Car) int64 {
	return c.price + c.tax
}

//객체 지향적인 구초제와 메소드를 바인딩하는 함수
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {
	//GO -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념 없음)
	// 객체 지향 -> 클래스(속성, 기능)

	// Go는 객체 지향 언어일까?
	// 인터페이스를 통한 다형성 지원, 구조체를 통해 클래스 형태의 코딩 가능하므로 (단, 클래스는 없다)
	// 객체 지향의 기본 개념을 가지고 있다.

	// 구조체와 메소드 연결을 통해 타 언어의 클래스 형식 처럼 사용 가능(객체 지향)

	// ex1
	bmw := Car{name: "520d", price: 50000000, color: "white", tax: 5000000}
	benz := Car{name: "220d", price: 60000000, color: "white", tax: 6000000}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex1 : ", benz, &benz)
	fmt.Printf("ex1 : %p\n", &bmw)
	fmt.Printf("ex1 : %p\n", &benz)

	fmt.Println("ex2 : ", Price(bmw), Price(benz)) //일반 함수에서 구조체를 호출한 방식 (객체 지향 X)
	fmt.Println("ex3 : ", bmw.Price(), benz.Price())

	fmt.Print("ex4 : ", &bmw == &benz)
}
