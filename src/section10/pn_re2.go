package main

import "fmt"

func runFunc() {
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred !")
	fmt.Println("실행되지않음")
}

func main() {
	//Go Recover 함수
	//에러 복구 가능
	//Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램이 종료되지 않는다!)

	//즉, 코드 흐름을 정상 상태로 복구하는 기능
	//Panic에서 설정한 메시지를 받아올 수 있다.

	//ex1
	runFunc()

	fmt.Println("Hello Golang!")
}
