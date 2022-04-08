package main

import (
	"fmt"
)

func main() {
	//채널 리소스 닫기 (패턴정리)
	//Close 주의 -> 닫힌 채널에 값 전송 시 패닉(예외) 발생
	//Range 채널 안에서 값을 꺼낸다. (순회), 채널 닫아야(close) 반복문 종료됨-> 채널이 열려있고 값 전송하지 않으면 계속 대기!

	// ex1
	ch := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) //위 5번 채널에 값 전송 후 채널 닫
	}()

	// 자주 사용하는 패턴!
	for i := range ch { // 채널에서 값을 꺼내온다. 채널이 close될때까지
		fmt.Println("ex1 : ", i)
	}

}
