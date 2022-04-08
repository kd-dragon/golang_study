package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//고루틴
	//클로저(closure) 사용 예제
	runtime.GOMAXPROCS(1) //cpu core 1개만 사용

	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) // 반복문 클로저는 일반적으로 즉시 실행하나 But 고루틴 클로저는 가장 나중에 실행 된다. (반복문 종료 후!)
		// 자주 쓰이지는 않지만 개념적으로 이해하자
	}

	time.Sleep(3 * time.Second)
}
