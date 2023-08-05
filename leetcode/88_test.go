package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge88(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		got := merge88([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
		assert.Equal(t, got, []int{1, 2, 2, 3, 5, 6})
	})
	t.Run("case2", func(t *testing.T) {
		got := merge88([]int{1}, 1, []int{}, 0)
		assert.Equal(t, got, []int{1})
	})
	t.Run("case3", func(t *testing.T) {
		got := merge88([]int{0, 0}, 0, []int{2, 5, 6}, 3)
		assert.Equal(t, got, []int{2, 5, 6})
	})
}

func Benchmark_merge88(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge88([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}
