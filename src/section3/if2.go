package main

import "fmt"

func main() {
	// 제어문 if~else문

	var a int = 50
	b := 70 // main 함수 안에서만 사용할 짧은 선언

	//ex1
	if a >= 65 {
		fmt.Println("a는 65 이상")
	} else {
		fmt.Println("a는 65 이하")
	}
	if b <= 70 {
		fmt.Println("b는 70 이상")
	} else {
		fmt.Println("b는 70 이하")
	}

  
}
