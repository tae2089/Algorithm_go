package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		got := maxProfit([]int{7, 1, 5, 3, 6, 4})
		assert.Equal(t, 5, got)
	})
	t.Run("case2", func(t *testing.T) {
		got := maxProfit([]int{7, 6, 4, 3, 1})
		assert.Equal(t, 0, got)
	})
}
