package main

import (
	"fmt"
)

func checkType(arg interface{}) {
	switch arg.(type) {
	case Account:
	case *Account:
	default:
	}
}

func checkType(arg interface{}) {
	//arg.(type)  을 통해 현재 데이터형 반환

	// golang은 break 문 없음
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a integer : ", arg)
	case float32, float64:
		fmt.Println("This is a float : ", arg)
	case string:
		fmt.Println("This is a string: ", arg)
	case nil:
		fmt.Println("This is a nil: ", arg)
	default:
		fmt.Println("What is this type ?", arg)
	}
}

func main() {
	/*
	   실제 타입 검사 switch 주로 사용
	   빈 인터페이스는 어떤 자료형도 전달 받으므로, 타입 체크를 통해 형 변환 후 사용한다.
	*/

	//ex1
	checkType(true)
	checkType(1)
	checkType(22.543)
	checkType(nil)
	checkType("Hello Golang!")
}
