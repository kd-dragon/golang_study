// 루리웹(ruliweb.com) 웹 크롤링 실습
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	//외부 라이브러리
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	urlRoot = "http://ruliweb.com"
)

//첫번째 방문 메인페이지 대상으로 원하는 url을 파싱 후 반환하는 함수
func parseMainNodes(n *html.Node) bool {
	//찾고자하는 tag
	if n.DataAtom == atom.A && n.Parent != nil {
		return scrape.Attr(n.Parent, "class") == "row"
	}
	return false
}

//errCheck is a function checking error
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

//동기화를 위한 작업 그룹 선언
var wg sync.WaitGroup

//url 대상이 되는 페이지(서브페이지) 대상으로 원하는 내용을 파싱 후 반환
func scrapContents(url string, fn string) {

	//작업종료 알림
	defer wg.Done()

	res, err := http.Get(url)
	errCheck(err)

	//요청 Body 닫기
	defer res.Body.Close()

	root, err := html.Parse(res.Body)
	errCheck(err)

	// Response 데이터 (html) 의 원하는 부분 파싱
	matchNode := func(n *html.Node) bool {
		return n.DataAtom == atom.A && scrape.Attr(root, "class") == "deco"
	}

	//파일 스트림 생성-> 파일명, 옵션, 권한
	file, err := os.OpenFile("G:/golang_crawling/"+fn+".txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)

	defer file.Close()

	w := bufio.NewWriter(file)

	//matchNode 메소드롤 사용하여 올바른 순회 화면서 출력
	for _, g := range scrape.FindAll(root, matchNode) {
		fmt.Println("result : ", scrape.Text(g))

		//파싱데이터 -> 버퍼에 기록
		w.WriteString(scrape.Text(g) + "\r\n")
	}
	w.Flush()
}

func main() {
	// main get 방식 요청
	res, err := http.Get(urlRoot)
	errCheck(err)

	//요청 Body 닫기
	defer res.Body.Close()

	//응답데이터(html)
	root, err := html.Parse(res.Body)
	errCheck(err)

	//ParsMainNodes: 메소드를 크롤링(스크랩핑) 대상 URL 추출
	urlList := scrape.FindAll(root, parseMainNodes)

	for _, link := range urlList {
		//대상 URL 1차 출력
		//fmt.Println("check main link : ", link, idx)
		//fmt.Println("TargetURL : ", scrape.Attr(link, "href"))
		fileName := strings.Replace(scrape.Attr(link, "href"), "https://bbs.ruliweb.com/family/", "", 1)
		fmt.Println("fileName : ", fileName)
		//작업 대기열에 추가
		wg.Add(1)

		//URL 및 해당
		go scrapContents(scrape.Attr(link, "href"), fileName)
	}

	wg.Wait()
}
