//함수 기초(1)
package main

import "fmt"
import "strconv" // 형변환 라이브러리
import "reflect"

func init() {
	fmt.Println("init method !")
}

func helloGolang() {
	fmt.Println("ex1: Hello Golang!!")
}

func say_one(m string) {
	fmt.Println("ex2:", m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	/*
	   함수
	   선언: func 키워도러 선언
	   func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
	   func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재
	   func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
	   func 함수명() : 매개변수 없음, 반환 값 없음

	   # 타 언어와 달리 반환 값(return value) 다수 가능
	   함수 선언 위치는 어느 곳이든 가능하다.
	*/

	// ex1
	helloGolang()

	// ex2
	say_one("Hello World")

	// ex3
	result := sum(5, 5)
	fmt.Println("ex3 :", reflect.TypeOf(result))
	fmt.Println("ex3 :", reflect.TypeOf(strconv.Itoa(sum(5, 5))))
	fmt.Println("ex3 :", result)
	fmt.Println("ex3 :", strconv.Itoa(sum(5, 5)))
}
