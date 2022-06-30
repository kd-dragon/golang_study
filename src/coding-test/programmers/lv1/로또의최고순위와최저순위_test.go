package lv1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLottoBestOrWorst(t *testing.T) {

	t.Run("테스트1", func(t *testing.T) {
		//given
		lottos := []int{44, 1, 0, 0, 31, 25}
		win_nums := []int{31, 10, 45, 1, 6, 19}

		//when
		rslt := LottoBestOrWorst(lottos, win_nums)

		//then: [3, 5]
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.ElementsMatch(t, rslt, []int{3, 5})
	})

	t.Run("테스트2", func(t *testing.T) {
		//given
		lottos := []int{0, 0, 0, 0, 0, 0}
		win_nums := []int{38, 19, 20, 40, 15, 25}

		//when
		rslt := LottoBestOrWorst(lottos, win_nums)

		//then: [1, 6]
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.ElementsMatch(t, rslt, []int{1, 6})
	})

	t.Run("테스트3", func(t *testing.T) {
		//given
		lottos := []int{45, 4, 35, 20, 3, 9}
		win_nums := []int{20, 9, 3, 45, 4, 35}

		//when
		rslt := LottoBestOrWorst(lottos, win_nums)

		//then: [1, 1]
		fmt.Println(rslt)

		assert.NotNil(t, rslt)
		assert.ElementsMatch(t, rslt, []int{1, 1})
	})
}
