package list

import (
	"math"

	"github.com/anthoturc/go-skippy/internal/node"
)

// TODO: This should probably be a configuration value
const MaxHeight = 17

type skipList struct {
	head *node.SkipNode
	tail *node.SkipNode
	size uint
}

type SkipList interface {
	Insert(item int)

	Find(item int) int

	Contains(item int) bool

	Delete(item int)

	Size() uint
}

func NewSkipList() SkipList {
	return newSkipList()
}

func newSkipList() *skipList {
	head := node.NewSkipNode(math.MinInt, MaxHeight, MaxHeight)
	head.Height = MaxHeight

	tail := node.NewSkipNode(math.MaxInt, MaxHeight, MaxHeight)
	tail.Height = MaxHeight
	for i := 0; i < MaxHeight; i++ {
		head.Next[i] = tail
		tail.Next[i] = nil
	}
	return &skipList{
		head: head,
		tail: tail,
		size: 0,
	}
}
