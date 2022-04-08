package main

import (
	"fmt"
)

func rangeSum(n int, c chan int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	c <- sum
	//채널로 전송이므로 return 값은 없다.
}

func main() {
	//Channel

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(700, c)
	go rangeSum(500, c)

	//순서대로 데이터 수신(동기) : 채널에서 값 수신 완료 될때까지 대기

	result1 := <-c //채널로부터 data수신
	result2 := <-c
	result3 := <-c

	fmt.Println("result1 : ", result1)
	fmt.Println("result2 : ", result2)
	fmt.Println("result3 : ", result3)

}
