package main

import (
	"fmt"
)

func main() {
	//Panic & Recover
	//사용자가 에러발생 가능
	//Panic 함수는 호출 즉시, 해당 메소드를 즉시 중지시키고 defer 함수를 호출하고 자기자신을 호출한 곳으로 리턴
	//런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생시킬 때 중요!
	//문법적인 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

	fmt.Println("Start Main")
	panic("Error occurred : user stopped !") //방법1
	//log.Panic("Error occurred : user stopped !") //방법2

	fmt.Println("End Main") //실행 불가

}