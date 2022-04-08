package main

import (
	"fmt"
)

func multiply1(x, y int) (r int) {
	r = x * y
	return r
}

func sum1(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	//함수를 변수에 할당! 중요!

	// ex1 (slice assignment)
	f := []func(int, int) int{multiply1, sum1}
	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println("ex1 : ", a, f[0](10, 10))
	fmt.Println("ex1 : ", b, f[1](10, 10))

	// ex2 func variable
	var f1 func(int, int) int = multiply1
	f2 := sum1 // code에서 자주 쓰인다.
	fmt.Println("ex1 : ", f1(10, 10))
	fmt.Println("ex1 : ", f2(10, 10))

	// ex3 (Map assignment)
	m := map[string]func(int, int) int{
		"mul_func": multiply1,
		"sum_func": sum1,
	}

	fmt.Println("ex3 :", m["mul_func"](10, 10))
	fmt.Println("ex3 :", m["sum_func"](10, 10))

}
