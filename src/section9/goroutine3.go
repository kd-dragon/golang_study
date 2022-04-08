package main

import (
	"fmt"
	"math/rand" //random 함수 사용
	"runtime"   //os 정보, 명령어 접근
	"time"
)

func exe(name int) {
	r := rand.Intn(100) // 100미만의 랜덤 숫자 추출
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}

func main() {
	//고루틴
	//멀티 코어 (다중 cpu) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                        //현 시스템의 CPU 코어 개수 반환 후 설정
	fmt.Println("Current System CPU : ", runtime.GOMAXPROCS(0)) // 설정 값 출력

	// ex1
	fmt.Println("main Routine Start : ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i) //고루틴 100개 생성
	}

	time.Sleep(5 * time.Second)
	fmt.Println("main Routine End : ", time.Now())

}
