//데이터 타입: 문자열(1)

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*
	   문자열
	   큰따옴표 "", 백스쿼트 `` 사용
	   Golang: 문자 char 타입 존재하지 않음. -> rune(int32의 alias)을 사용하여 아스키코드값으로 표현
	   문자 : ''로 작성
	   자주 사용하는 escape : \\, \', \", \a (콘솔벨), \b (백스페이스), \f (쪽바꿈), \n (줄바꿈), \r (복귀), \t(탭) ...
	*/

	//ex1
	var (
		str1 string = "c:\\go_stury\\src\\" // c:\go_study\src\
		str2 string = `c:\go_stury\src\`    // escape 사용으로 인해 지저분함을 없애준다.
	)

	fmt.Println("str1 : ", str1)
	fmt.Println("str2 : ", str2)

	//ex2
	var str3 string = "Hello, world"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00" //유니코드

	fmt.Println("str3 : ", str3)
	fmt.Println("str4 : ", str4)
	fmt.Println("str5 : ", str5)

	//ex3
	//GO에서는 문자열의 길이(실제 길이)와 크기(바이트 수)를 구하는 것을 명확하게 지원해준다.
	//길이(바이트 4)
	fmt.Println("ex3 eng : ", len(str3)) //영어는 글자당 1바이트
	fmt.Println("ex3 kor : ", len(str4)) //한글은 글자당 3바이트

	//ex4
	//길이(실제길이)
	fmt.Println("ex4 eng : ", utf8.RuneCountInString(str3))
	fmt.Println("ex4 kor : ", utf8.RuneCountInString(str4))
	fmt.Println("ex4 kor : ", len([]rune(str4))) // len으로 실제 길이 구하는법 (일반적으로 RuneCountInString을 많이쓴다)

}
