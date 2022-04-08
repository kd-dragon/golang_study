package main

import "fmt"

func main() {
	// slice copy

	// copy(target, origin)
	// 반드시 make로 공간을 할당 후 복사해야한다.
	// 복사된 슬라이스 값 변경해도 원본에는 영향 없음 !

	// ex1
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 5)
	s3 := []int{}

	copy(s2, s1) //s2 에 s1을 copy
	copy(s3, s1)

	fmt.Println("ex1 : ", s2) //cap 5이므로 5까지만 카피됨
	fmt.Println("ex1 : ", s3) //복사가 안됨. (공간이 할당되지 않았으므로)

	// ex2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)
	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)

	fmt.Println()

	//ex3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] // 주의! 부분적 슬라이스 추출은 참조 -> 원본 값 변경된다!

	d[1] = 7

	fmt.Println("ex3 : ", c)
	fmt.Println("ex3 : ", d)

	fmt.Println()

	//ex4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 0부터 5까지 슬라이싱하고(참조값 복사) 7만큼의 용량으로 복사해라.
	f[1] = 999
	fmt.Println("ex4 : ", f, len(f), cap(f))
	fmt.Println("ex4 : ", e)
}
