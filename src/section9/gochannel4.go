package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Channel

	runtime.GOMAXPROCS(1)

	// ex1 (비동기: 버퍼 사용)
	//발신자 -> 가득 차면 대기 / 비어있으면 작동
	//수신자 -> 비어있으면 대기 / 가득 차면 작동

	ch := make(chan bool, 2)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go : ", i)
		}

	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}
}
