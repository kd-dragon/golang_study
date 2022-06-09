package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Go 에러처리 기초 (1)
	// 소프트웨어의 품질 향상시키는 데 가장 중요한 것! -> 유형코드 및 에러 정보 등을 남기는 행위

	// Go에서는 기본적으로 error 패키지에서 error 인터페이스를 제공
	// Go언어 자체는 error처리가 없다.

	// type error interface { Error() string }
	//즉 , Error 메서드를 구현하면 사용자 정의 에러 처리 제작 가능
	//기본적으로 메서드마다 리턴 타입 2개. (데이터 리턴값, 에러값)
	// 방법1) Fatal(프로그램 종료) 메서드
	// 방법2) 주로 Errorf(에러 타입 리턴) 메서드

	//기본적인 메서드 에러처리 예제
	f, err := os.Open("unnamedfile") // error 발생

	//error가 nil이오면 정상. nil이 아니면 에러발생한 것 .
	if err != nil {
		log.Fatal(err.Error()) //방법1 -> Fatal은 프로그램 종료가 되버린다.
		//log.Fatal(err) // 동일한 결과 출력
	}
	fmt.Println("==========================")
	fmt.Println("ex1 : ", f.Name())
}
