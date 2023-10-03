package set

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipListInit(t *testing.T) {
	sl := newSkipListSet[int]()

	assert.NotNil(t, sl.head)
	assert.NotNil(t, sl.tail)
	assert.Equal(t, MaxHeight, sl.head.Height)
	assert.Equal(t, MaxHeight, sl.tail.Height)

	for i := 0; i < MaxHeight; i++ {
		assert.Equal(t, sl.tail, sl.head.Next[i])
	}
}

func TestEmptySkipLis(t *testing.T) {
	sl := newSkipListSet[int]()

	assert.False(t, sl.Contains(2))
	assert.False(t, sl.Contains(1))
	assert.False(t, sl.Contains(math.MaxInt))
	assert.False(t, sl.Contains(math.MinInt))
}

func TestInsertSkipList(t *testing.T) {
	sl := newSkipListSet[int]()

	for i := 0; i < 10; i++ {
		sl.Insert(i)
	}

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("skip-list-contains-%d", i), func(t *testing.T) {
			assert.True(t, sl.Contains(i))
		})

	}

	for i := 10; i < 20; i++ {
		t.Run(fmt.Sprintf("skip-list-does-not-contain-%d", i), func(t *testing.T) {
			assert.False(t, sl.Contains(i))
		})
	}
}

func TestInsertSkipListDuplicates(t *testing.T) {
	sl := newSkipListSet[int]()

	n := 10
	for i := 0; i < n; i++ {
		sl.Insert(i)
		sl.Insert(i)
	}

	for i := 0; i < n; i++ {
		assert.True(t, sl.Contains(i))
	}

	assert.Equal(t, uint(n), sl.Size())
}

func TestDeleteSkipList(t *testing.T) {

	sl := newSkipListSet[int]()

	n := 10

	for i := 0; i < n; i++ {
		sl.Insert(i)
		assert.Equal(t, uint(i+1), sl.Size())
	}
	assert.Equal(t, uint(n), sl.Size())

	for i := n - 1; i >= 0; i-- {
		assert.True(t, sl.Contains(i))
		sl.Delete(i)
		assert.False(t, sl.Contains(i))
		assert.Equal(t, uint(i), sl.Size())
	}
}
