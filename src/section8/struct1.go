package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

//Receiver
func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//구조체
	//Go의 필드드의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다. (리시버 형식으로 연결해서 사용)
	//(상속,객체,클래스 개념없음)

	// ex1
	kim := Account{number: "235-901", balance: 10000000, interest: 0.015}
	lee := Account{number: "235-902", balance: 12000000} //기본값 0으로 초기화
	park := Account{number: "235-903", interest: 0.025}
	cho := Account{"235-904", 15000000, 0.03}

	fmt.Println("ex1: ", kim)
	fmt.Println("ex1: ", lee)
	fmt.Println("ex1: ", park)
	fmt.Println("ex1: ", cho)

	// ex2
	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))

}
