package lv1

import "strings"

// https://programmers.co.kr/learn/courses/30/lessons/92334

func ReportingResult(id_list []string, report []string, k int) []int {

	reportMap := make(map[string]int)
	rsltMap := make(map[string]int)
	rslt := make([]int, 0)

	uniqReport := makeSliceUnique(report)
	for _, v := range uniqReport {
		recv := strings.Split(v, " ")[1]
		reportMap[recv]++
	}

	for _, v := range uniqReport {
		send := strings.Split(v, " ")[0]
		recv := strings.Split(v, " ")[1]
		if reportMap[recv] >= k {
			rsltMap[send]++
		}
	}

	for _, v := range id_list {
		rslt = append(rslt, rsltMap[v])
	}

	return rslt
}

func makeSliceUnique(s []string) []string {
	keys := make(map[string]struct{})
	res := make([]string, 0)

	for _, v := range s {
		if _, ok := keys[v]; ok {
			continue
		}
		keys[v] = struct{}{}
		res = append(res, v)
	}
	return res
}
