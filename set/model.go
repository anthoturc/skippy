package set

import (
	"math"

	"github.com/anthoturc/skippy/internal/node"
)

// TODO: This should probably be a configuration value
const MaxHeight = 17

type skipListSet struct {
	head *node.SkipNode
	tail *node.SkipNode
	size uint
}

type SkipListSet interface {
	Insert(key int)

	Contains(key int) bool

	Delete(key int)

	Size() uint
}

func NewSkipListSet() SkipListSet {
	return newSkipListSet()
}

func newSkipListSet() *skipListSet {
	head := node.NewSkipNode(math.MinInt, MaxHeight, MaxHeight)
	head.Height = MaxHeight

	tail := node.NewSkipNode(math.MaxInt, MaxHeight, MaxHeight)
	tail.Height = MaxHeight
	for i := 0; i < MaxHeight; i++ {
		head.Next[i] = tail
		tail.Next[i] = nil
	}
	return &skipListSet{
		head: head,
		tail: tail,
		size: 0,
	}
}
