//사용자 패키지 작성 및 문서화 예제
package main

import (
	"fmt"
	oper "section12/arithmetic"
)

//alias 사용(패키지 중복 또는 약자로 사용)
//"section12/arithmetic" -> alias 미사용시 그대로 입력

func main() {
	//사용자 패키지 작성 & Go 문서화

	//기준은 GOPATH !!
	//패키지 폴더명(디렉토리명)
	//패키지 파일의 package 이름으로 사용한다. -> 길면 alias 사용
	//package main을 제외하고 package 문서에 등록
	//지금까지 우리는 패키지를 사용해 왔다.
	//기본적으로 GOROOT의 패키지 (pkg) 폴더를 검색하고 없으면 GOPATH의 패키지(src/pkg)를 검색합니다.
	//go install 명령어 실행의 경우 -> GOPATH/pkg 에 등록됨 (문서화)
	//godoc -http=:8080(임의의포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인 (패키지, 타입, 메소드) 주석

	//패키지 사용 예제(사칙연산)
	nums := oper.Numbers{X: 100, Y: 10}
	fmt.Println("Package Used(1) : ", nums.Plus())
	fmt.Println("Package Used(2) : ", nums.Minus())
	fmt.Println("Package Used(3) : ", nums.Multi())
	fmt.Println("Package Used(4) : ", nums.Divide())
	fmt.Println("Package Used(5) : ", nums.SquarePlus())
	fmt.Println("Package Used(6) : ", nums.SquareMinus())

}
