package main

import "fmt"

func main() {
	// ex1
	i := 7
	p := &i

	fmt.Println("ex1 : ", i, &i, p, *p)
	fmt.Printf("ex1 : %d %p %p %d", i, &i, p, *p)
	fmt.Println()
	*p++ //역참조해서 숫자1을 증가시킴
	fmt.Printf("ex1 : %d %p %p %d\n", i, &i, p, *p)
	*p = 77777 //포인터 변수 역 참조 값 변경
	fmt.Printf("ex1 : %d %p %p %d\n", i, &i, p, *p)
	i = 77
	fmt.Printf("ex1 : %d %p %p %d\n", i, &i, p, *p)
}
