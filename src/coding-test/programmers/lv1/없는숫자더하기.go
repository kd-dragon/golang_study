package lv1

//https://programmers.co.kr/learn/courses/30/lessons/86051

func AddNotExistNum(numbers []int) int {
	var sum int
	max := 45

	for _, val := range numbers {
		sum = sum + val
	}

	return max - sum
}
