package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 뮤텍스: 상호 배제 -> Thread(고루틴)들이 서로 running time에 서로 영향을 주지 않도록 하는 기술.(무결성 보장)

	//RWMutex : 쓰기 Lock -> Write 중에는 다른 곳에서 이전 값을 읽으면 안됨. (읽기, 쓰기 Lock 전부 방지!)
	//RMutex: 읽기 Lock -> Read 중에 값 변경 방지! 쓰기 Lock만 방지
	//반드시 Lock, Unlock 짝이 맞아야한다.

	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			//쓰기 뮤텍스 잠금
			mutex.Lock()
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
			mutex.RUnlock()
		}
	}()
	// RLock끼리는 서로 방지하지 않음.
	go func() {
		for i := 1; i <= 10; i++ {
			//읽기 뮤텍스 잠금
			mutex.RLock()
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
			fmt.Println("Read2 : testssss")
			mutex.RUnlock()
		}
	}()

	time.Sleep(12 * time.Second)
}
