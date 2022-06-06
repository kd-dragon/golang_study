package main

import "fmt"
import "time"

func main() {
	// channel 셀렉트 구문
	// 채널에 값이 수싲ㄴ되면 해당 case 문 실행
	// 일회성 구문이므로 for(반복문)안에서 수행한다.
	//default 구문 처리 주의.

	//ex1
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Hi Golang!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
				// default 되도록 사용하지말자 (값이 아니면 무조건 default 찍히므로)
			}
		}
	}()

	time.Sleep(7 * time.Second)

}
