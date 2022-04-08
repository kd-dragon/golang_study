//패키지(1)
package main

import (
	"fmt"
	"os"
)

//선언 방법
// import "fmt"
// import "os"

//선언 방법2 (좀 더 자주 쓰임)

func main() {
	//패키지: 코드 구조화 (모듈) 및 재사용
	//응집도 느슨하게, 결합도는 높게해야 유지보수 쉽고 가독성 높고 재사용성이 높은 clean 코드가 된다.
	//GO: 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그래밍을 작성할 것을 권고
	//패키지 이름  = 디렉토리 이름
	//같은 패키지 내  -> 소스 파일들은 디렉토리명을 패키지 명으로 사용한다.
	//네이밍 규칙 : 소문자 -> private, 대문자 -> public

	//** Main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 X, 프로그래머의 시작점 start Point로 인식

	var name string
	fmt.Println("이름은? : ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi! %s\n", name)
}
