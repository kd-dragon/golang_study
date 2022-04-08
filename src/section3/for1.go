package main

import "fmt"

func main() {

	/*
	   반복문 - For
	   Go에서 유일하게 제공되는 반복문
	   다양한 사용 패턴 숙지
	*/
	// ex1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}
	// ex2 무한 루프 패턴.
	/*
	   for {
	     fmt.println("ex2 : Hello, Golang!")
	     fmt.println("ex2 : Infinite loop!")
	   }
	*/
	// ex3 (Range)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
	// key value 형식이 아니라 첫번째는 인덱스, 두번째는 value
	for name := range loc {
		fmt.Println("ex3 : ", name) // index가 찍힌다.
	}
	// 해결방안 -> _ 언더바 사용 (생략한다는 표현: iota(열거형) 변수 배울때 사용했었음)
	for _, name := range loc {
		fmt.Println("ex3 : ", name) // index가 찍힌다.
	}

  
}
