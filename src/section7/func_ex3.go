// 재귀함수
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi !")
	prtHello(n - 1)
}

func main() {
	/*
	   *함수 고급
	   재귀 함수 (Recursion)
	   장점: 프로그램이 보기 쉽고 코드 간결 오류 수정 용이
	   단점: 코드를 이해하기 어렵고, 기억 공간 많이 사용, 무한루프 위험있음
	*/

	// ex1
	x := fact(5)
	fmt.Println(x)

	// ex2
	prtHello(5)
}
