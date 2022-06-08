package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	//고루틴 동기화 객체
	//동기화 상태(조건) 메소드 사용
	//Wait, notify, notifyAll : 기타 언어에서 사용하는 메소드 명
	//Wait, Signal, Broadcas : Go언어

	//고루틴을 Wait시키면 Signal이나 Broadcas가 올때까지 일시정지시킨다.

	//시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) //비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- n
			fmt.Println("Goroutine waiting : ", n)
			condition.Wait() //고루틴 대기 (멈춤)
			fmt.Println("Wating End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		// <-c
		fmt.Println("received :", <-c)
	}

	//하나씩 깨우는 예제
	/*
		for i := 0; i < 5; i++ {
			mutex.Lock()
			fmt.Println("Awaken Goroutine(Signal) :", i)
			condition.Signal() //한개씩 깨움 (모든 고루틴 생성 후)
			mutex.Unlock()
		}
	*/
	//전체 한번에 깨우는 예제
	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)

}
