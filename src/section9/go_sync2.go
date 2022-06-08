package main

import (
	"fmt"
	"runtime"
	"sync"
)

//구조체 선언(공유 데이터)
type count struct {
	num   int
	mutex sync.Mutex //mutex 구조체에 선언
}

func (c *count) increment() {
	c.mutex.Lock()
	c.num++
	c.mutex.Unlock()
}

//구조체는 포인터형으로 받는 경우가 많다. (read만 하더라도)
func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 기초 (1)
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요하다.
	//뮤텍스(Mutex) : 여러 고루틴에서 작업하는 공유데이터 보호
	//sync.Mutex 선언 후 Lock, Unlock 사용!

	//ex1 (동기화 사용한 경우)
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
