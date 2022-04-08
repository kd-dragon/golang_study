package main

import "fmt"

func main() {
	// 제어문(조건문)
	/* IF문 주요 규칙 :
	   1.반드시 Boolean 검사  (1, 0 사용불가 : 자동 형 변환 불가하다)
	   2.소괄호 사용 불가
	*/

	var a int = 20
	b := 20

	// ex1
	if a >= 15 {
		fmt.Println("15이상")
	}

	// ex2
	if b >= 25 {
		fmt.Println("25이상")
	}

	// ex3
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

	// 자주 발생하는 에러 1 (Go 언어는 끝에 세미콜론을 붙이므로)
	/*
	   if b > = 25
	   {

	   }
	*/

	// 자주 발생하는 에러 2 (조건 1개여도 괄호 생략 불가)
	/*
	   if b >= 25
	     fmt.Println("25이상")
	*/

	// 자주 발생하는 에러 3 (Boolean 대신 1,0 사용 불가)
	/*
	  if c:=1; c {
	    fmt.Println("True")
	  }
	*/

}
