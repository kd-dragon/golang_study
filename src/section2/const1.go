// 상수 1
package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 후 값 변경 금지, 고정 된 값 관리용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getFunc() 불가: 함수가 매번 같은 값을 리턴해준다는 보장이 없으므로
	const e = 35.6
	const f = false

	/*
	   에러 발생
	   const g string
	   g = "test3"
	*/

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)
}
