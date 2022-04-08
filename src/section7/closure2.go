package main

import "fmt"

func main() {
	cnt := increaseCnt()
	/*
	   우리가 알고있는 기본적인 함수라면 매번 호출시마다 같은 결과를 반환해야한다. (매번 1이 찍혀야한다.)
	   그러나 클로저는 지역변수 n의 메모리 주소를 캡쳐해서 사용하므로 누적 카운팅이 가능하다.
	*/
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())

}

func increaseCnt() func() int {
	n := 0 // 지역변수(메모리의 주소가 캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
