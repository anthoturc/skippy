package list

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipListInit(t *testing.T) {
	sl := newSkipList()

	assert.NotNil(t, sl.head)
	assert.NotNil(t, sl.tail)
	assert.Equal(t, MaxHeight, sl.head.Height)
	assert.Equal(t, MaxHeight, sl.tail.Height)

	for i := 0; i < MaxHeight; i++ {
		assert.Equal(t, sl.tail, sl.head.Next[i])
	}
}

func TestEmptySkipLis(t *testing.T) {
	sl := newSkipList()

	assert.False(t, sl.Contains(2))
	assert.False(t, sl.Contains(1))
	assert.False(t, sl.Contains(math.MaxInt))
	assert.False(t, sl.Contains(math.MinInt))
}
