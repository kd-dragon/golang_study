package main

import (
	"errors"
	"fmt"
)

func main() {
	//에러 처리

	//errors 패키지의 new 메소드를 활용한 에러 생성
	//Errorf보다 더 많이 사용합니다!

	var err1 error = errors.New("Error occurred - 1")
	err2 := errors.New("Error occurred - 2")

	fmt.Println("err1 : ", err1.Error())
	fmt.Println("err2 : ", err2.Error())
}
