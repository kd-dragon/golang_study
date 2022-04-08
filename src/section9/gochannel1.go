package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ---> ", time.Now())
	v <- 1 //1을 채널로 전송 (송신)
}

func work2(v chan int) {
	fmt.Println("work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 : E ---> ", time.Now())
	v <- 2 //2을 채널로 전송 (송신)
}

func main() {
	//채널(Channel)
	//고루틴간의 상호 정보 (데이터) 교환 및 실행 흐름 동기화 위해서 사용
	//실행 흐름 제어 가능(동기, 비동기) --> 일반 변수로 선언 후 사용 가능
	//데이터 전달 자료형 선언 후 사용 (지정된 타입만 주고 받을 수 있음)

	//인터페이스(interface{}) 전달을 통해서 자료형 상관없이 전송 및 수신 가능
	//값 전달 (복사후 bool, int 등), 포인터(슬라이스, 맵) 등을 전달시에는 주의! > 동기화 사용(Mutex)

	//멀티프로세싱 처리에서 교착 상태(경쟁) 주의

	// <- , -> ( 채널 <- 데이터 : 송신) , ( <- 채널 :수신)

	// ex1
	fmt.Println("Main : S ---> ", time.Now())

	//var c chan int
	//c = make(chan int)
	v := make(chan int) // int형 채널 짧은 선언

	go work1(v)
	go work2(v)

	<-v
	<-v

	//채널은 동기식이므로 <-v 로 값을 받아오기전까지 메인함수가 끝나지않는다.
	fmt.Println("Main : E ---> ", time.Now())

}
