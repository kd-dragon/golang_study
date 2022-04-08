package main

import "fmt"

func a() {
	defer end(start("b")) //중첩 함수 주의! (이렇게 사용하지말고 왠만하면 함수 하나로만 쓰자)
	fmt.Println("in a func")
}

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func main() {
	// ex1
	a()
	/*
			   예상 순서는 다음과 같다.
			   in a func -> start : b -> end : b
			   그러나 실제 동작은
			   start : b -> in a func -> end : b

			   # 원인: defer 예약어는 바로 뒤에 있는 func에만 적용되므로 end함수에만 적용되었기 때문이다.
		     따라서 end() 매개변수인 start 함수를 먼저 실행해버린다.
	*/
}
