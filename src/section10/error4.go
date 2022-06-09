package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(n int) (string, error) { //메소드 리턴 값 error 타입 중요!
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	return "", errors.New("0을 입력했습니다. 에러 발생 ! ")
}

func main() {
	// Go 에러처리 기초 (1)
	// Errorf를 이용한 에러 처리

	a, err := notZero(1)

	//error가 nil이오면 정상. nil이 아니면 에러발생한 것 .
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ex1 : ", a)

	b, err := notZero(0)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Fatal 이후 프로그램 종료 후 실행 중지
	fmt.Println("ex2 : ", b)
	fmt.Println("End Error handling !")

}
