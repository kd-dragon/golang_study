package lv1

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
