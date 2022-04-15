package main

import "fmt"
import "reflect"

func main() {
	/*
	   Type assertion (타입 변환)

	   실행 (런타임)시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야한다.
	   이때 인터페이스.(타입) 으로 형변환한다.
	   ex) interfaceVal.(type)
	*/

	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64)
	//최초 데이터 타입으로만 변환 가능

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	fmt.Println()

	//ex2 (저장 ) //실제로 정말 많이쓰는 패턴
	if v, ok := a.(int); ok {
		fmt.Println("ex2 : ", v, ok)
	}
}
