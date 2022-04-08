package main

// 함수 심화 (1)

import (
	"fmt"
)

// ...int -> int형으로 몇개까지든 받을 수 있음을 명시
func multiplyEX(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sumEX(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2: ", value)
	}
}

func main() {
	/*
	   *함수 고급
	   가변 인자 실습 ( 매개 변수 개수가 동적으로 변할 때 - 정해져있지않은 경우 )
	*/

	// ex1
	x := multiplyEX(1, 2, 3)
	y := sumEX(1, 2, 3, 4, 5)
	fmt.Println("ex1 :", x)
	fmt.Println("ex1 :", y)
	fmt.Println()

	// ex2 문자열
	prtWord("a", "apple", "test", "golang", "good")

	// ex3 slice, array
	a := []int{1, 2, 3, 4, 5}
	m := multiplyEX(a...) // 그냥 a 가 아니라 a... 을 찍어주면 된다.
	n := sumEX(a...)
	fmt.Println(m)
	fmt.Println(n)

}
