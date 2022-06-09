package main

import (
	"errors"
	"fmt"
	"math"
)

//f의 i제곱 구하는 함수
func Power(f float64, i float64) (float64, error) {
	if f == 0 {
		return 0, errors.New("0은 사용 불가능 합니다.")
	}
	return math.Pow(f, i), nil
}

func main() {
	//에러 처리 고급
	//Go error 패키지 new 메소드 사용 -> 사용자 에러 처리 생성

	//ex1
	if f, err := Power(7, 3); err != nil {
		fmt.Println("Error message : %s \n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}

	//ex2
	if f, err := Power(0, 3); err != nil {
		fmt.Printf("Error message : %s \n", err.Error())
	} else {
		fmt.Println("ex1 : ", f)
	}
}
