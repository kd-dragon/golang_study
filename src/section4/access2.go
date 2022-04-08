package main

//패키지 접근제어(2)

import (
	"fmt"
	checkUp "section4/lib"
	_ "section4/lib2"
	//사용하진 않지만 지우지말고 남기려면 빈 식별자를 사용한다.
)

func main() {
	//패키지 접근 제어
	// 별칭 사용
	// 빈 식별자 사용 ( _ )

	fmt.Println("10보다 큰 수? : ", checkUp.CheckNum(11))
	//fmt.Println("100보다 큰 수? : ", checkUp2.CheckNum(100))
}
