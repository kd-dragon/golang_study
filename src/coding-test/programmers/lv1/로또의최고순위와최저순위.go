package lv1

// https://programmers.co.kr/learn/courses/30/lessons/77484

func LottoBestOrWorst(lottos []int, win_nums []int) []int {

	rules := []int{6, 6, 5, 4, 3, 2, 1}
	winCnt := 0
	zeroCnt := 0
	maxScore := 0

	//이긴 번호 갯수
	for _, v := range win_nums {
		if contains(lottos, v) {
			winCnt++
		}
	}

	// 알아볼 수 없는 번호 갯수
	for _, v := range lottos {
		if v == 0 {
			zeroCnt++
		}
	}

	maxScore = winCnt + zeroCnt

	if maxScore > 6 {
		maxScore = 6
	}

	//최고 순위, 최저 순위
	return []int{rules[maxScore], rules[winCnt]}
}

func contains(nums []int, searchNum int) bool {
	for _, n := range nums {
		if n == searchNum {
			return true
		}
	}
	return false
}
