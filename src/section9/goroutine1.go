package main

import (
	"fmt"
	"time"
)

/*
  고루틴
  타 언어의 쓰레드(Thread)와 비슷한 기능
  생성 방법 매우 간단, 리소스 매우 적게 사용 (Kbyte 단위)
  즉, 수많은 고루틴 동시 생성 실행 가능

  비동기적 함수 루틴 실행(매우 적은 용량 차지) --> 채널을 통한 통신 가능
  공유메모리 사용 시에 정확한 동기화 코딩 필요!!!

  // 싱글 루틴에 비해 항상 빠른 처리 결과는 아니다.

  *멀티쓰레드의 장점과 단점
  장점: 응답성 향상, 자원 공유를 효율적으로 활용, 작업이 분리되어 코드 간결
  단점: 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스의 사이드 이펙트, 성능 저하, 동기화 문제발생가능
  결론- 정확하게 알고 사용해야한다.

*/

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func End", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func End", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func End", time.Now())
}

func main() {
	exe1() // 가장 먼저 실행

	fmt.Println("Main Routine Start", time.Now())
	go exe2() // go 루틴은 이 한 줄이면 끝난다. (java처럼 runnable 상속 등 필요없음)
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine End", time.Now())

}
