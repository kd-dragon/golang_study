package main

import "fmt"

type cnt int

func main() {

	// 기본 자료형 사용자 정의 타입
	// ex1 (잘 사용하지 않음, 헷갈려서)
	a := cnt(5)
	fmt.Println("ex1 : ", a)

	// ex2
	var b cnt = 15
	fmt.Println("ex1 : ", b)

	// ex3
	// testCoverT(b) // 예외 발생(int형인데 왜 cnt형을 넣어?)
	testCoverT(int(b)) // int로 형변환 가능
	testCoverD(b)

}

func testCoverT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

// 구조체 cnt를 만들었으므로 파라미터로 기재
func testCoverD(i cnt) {
	fmt.Println("ex4 : (Custom Type) : ", i)
}
