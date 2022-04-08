package main

import "fmt"

func init() {
	fmt.Println("pointer3 init method !")
}

//reference call
func rptc(n *int) {
	*n = 77
}

//value call
func vptc(n int) {
	n = 77
}

func main() {
	// 함수, 메서드 호출 시에 매개변수 값을 복사 전달 -> 함수, 메서드 내에서는 원본 값 변경 불가능!!!
	// 원본 값 변경을 위해서는 포인터로 전달
	// 특히 크키가 큰 배열의 경우 값 복사시 자원 소모 크므로 포인터전다롤 해결(슬라이스, Map 참조 전달)

	//ex1
	var a int = 10
	var b int = 10
	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)
	fmt.Println()

	rptc(&a)
	vptc(b)

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)

}
