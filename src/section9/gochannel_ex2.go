package main

import (
	"fmt"
)

//수신할때쓰는 channel을 리턴값으로 정의
func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := 1; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()

	return tot
}

func main() {
	//channel
	//채널 또한 함수의 반환 값으로 사용 가능

	//ex1
	c := sum(100)

	fmt.Println("ex1 : ", <-c)
}
