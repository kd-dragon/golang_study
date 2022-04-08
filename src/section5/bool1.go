//데이터 타입: Bool(1)
package main

import "fmt"

func main() {
	//Bool (Bollean): true, false
	// 조건부 논리 연산자와 함께 사용 : !, ||, &&
	// 암묵적 형변환X : 0, Nil -> False 변환 없음!!! 무조건 false로 써줘야함

	//ex1
	var b1 bool = true
	var b2 bool = false
	b3 := true
	//b4 := 1

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex2 : ", b2)
	fmt.Println("ex3 : ", b3)

	fmt.Println("ex4 : ", b3 == b1)
	fmt.Println("ex5 : ", b1 && b3)
	fmt.Println("ex6 : ", b1 || b2)
	fmt.Println("ex7 : ", !b1)

	/* //암묵적 형변환 X
	if b4 {
		fmt.Println("ex8 : true")
	}
	*/

}
