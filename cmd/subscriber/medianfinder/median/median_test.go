package median_test

import (
	"simple-pub-sub/cmd/subscriber/medianfinder/median"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian_FindsMedian(t *testing.T) {
	var nums = []struct {
		In  []int
		Exp int
	}{
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{30, 10, 20}, 20},
		{[]int{4, 6}, 5},
	}

	for _, n := range nums {
		out := median.Find(n.In)
		assert.Equal(t, n.Exp, out)
	}
}
