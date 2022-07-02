package lv1

// https://programmers.co.kr/learn/courses/30/lessons/76501

func SumAbsolutes(absolutes []int, signs []bool) int {

	var sum int

	for idx, sign := range signs {
		if sign {
			sum += absolutes[idx]
			continue
		}
		sum -= absolutes[idx]
	}

	return sum
}
