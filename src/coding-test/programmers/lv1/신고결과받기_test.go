package lv1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportingResult(t *testing.T) {

	t.Run("테스트1", func(t *testing.T) {
		//given
		id_list := []string{"muzi", "frodo", "apeach", "neo"}
		report := []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}
		k := 2

		//when
		rslt := ReportingResult(id_list, report, k)

		//then: [2,1,1,0]
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.ElementsMatch(t, rslt, []int{2, 1, 1, 0})
	})

	t.Run("테스트2", func(t *testing.T) {
		//given
		id_list := []string{"con", "ryan"}
		report := []string{"ryan con", "ryan con", "ryan con", "ryan con"}
		k := 3

		//when
		rslt := ReportingResult(id_list, report, k)

		//then: [0,0]
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.ElementsMatch(t, rslt, []int{0, 0})
	})
}
