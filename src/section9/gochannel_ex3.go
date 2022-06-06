package main

import (
	"fmt"
)

func receiveOnly(cnt int) <-chan int {

	sum := 0
	tot := make(chan int)

	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 888
		close(tot) //채널 문이 닫히고 값만 가지고 있게된다.
	}()

	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a += i
		}
		tot <- a
	}()

	return tot
}

func main() {
	//channel

	//ex1
	c := receiveOnly(1000) //채널 반환
	output := total(c)     //채널 전달 후 반환

	fmt.Println("ex1 : ", <-output)
}
