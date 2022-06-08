package main

import (
	"fmt"
	"runtime"
)

//구조체 선언(공유 데이터)
type count struct {
	num int
}

func (c *count) increment() {
	c.num++
	// +1이 반영되기전 다른 고루틴이 읽어가면 문제 발생
}

//구조체는 포인터형으로 받는 경우가 많다. (read만 하더라도)
func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 기초 (1)
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요하다.

	//ex1 (동기화 사용하지 않은 경우)
	// 시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() //Cpu 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}
	c.result()

}
