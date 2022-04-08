package main

import "fmt"

func main() {
	/*
	   <포인터>
	   변수의 지역성 (외부에서 사용한 변수를 사용할 수 있다.)
	   파이썬, 자바는 포인터를 지원하지 않는다. (Native 코드(내부적인 코드)에서는 사용하고 있다.)

	   포인터지원X 언어(파이썬, C#, JAVA 등)

	   GO는 제한적으로만 포인터 기능 제공
	   - 주소의 값은 직접 변경 불가능 (잘못된 코딩으로 인한 버그 방지)
	   - *(에스터리스크) 로 사용
	   - nil 로 초기화 (nil 은 0이 아니다)
	*/

	// ex1

	var a *int            //Pattern1
	var b *int = new(int) //Pattern2

	fmt.Println(a)
	fmt.Println(b) //new 할당하면 주소값을 가진다.

	i := 7

	fmt.Println("ex1 : ", 1, &i)
	fmt.Println()
	a = &i //주소 값 전달
	b = &i
	*a = 77
	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", &a)
	fmt.Println("ex1 : ", *a) //역참조
	fmt.Println()
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b) //역참조

	var c = &i
	d := &i

	*d = 10

	fmt.Println("ex2 : ", *c)
	fmt.Println("ex2 : ", *d)

}
