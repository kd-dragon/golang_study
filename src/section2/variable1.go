//변수1
package main

import "fmt" // 선언해놓고 사용안하면 저장시 자동으로 제거된다.

func main_() {
	//기본 초기화
	//정수타입 : 0, 실수: 0.0, 문자열: "", Boolean : true, false
	//변수명 : 숫자 첫글자 X, 대소문자 구분 O, 문자 숫자 밑줄 특수기호 사용 가능(한글도 가능)
	//변수 및 상수 : 함수 내외 사용 가능
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"
	var k = 4.74        // 선언 동시 초기화
	var l = "Hi Seoul!" // 선언 동시 초기화
	var m = true

	a = 4
	b = "Hi! Ray Kim!"
	e = 77
	// GO에서는 변수 선언을 하고 사용하지 않으면 에러가 발생합니다.
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

}
