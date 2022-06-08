package main

import (
	"fmt"
	_ "runtime"
	_ "sync"
	"time"
)

func main() {
	// 뮤텍스: 상호 배제 -> Thread(고루틴)들이 서로 running time에 서로 영향을 주지 않도록 하는 기술.(무결성 보장)

	data := 0

	go func() {
		for i := 1; i <= 10; i++ {
			data += 1
			fmt.Println("Write : ", data)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read1 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println("Read2 : ", data)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
}
