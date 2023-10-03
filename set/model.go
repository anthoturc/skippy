package set

import (
	"math/rand"
	"time"

	"github.com/anthoturc/skippy/internal/node"
	"golang.org/x/exp/constraints"
)

// TODO: This should probably be a configuration value
const MaxHeight = 17

type skipListSet[K constraints.Ordered] struct {
	head *node.SkipNode[K]
	tail *node.SkipNode[K]
	size uint
	rng  *rand.Rand
}

type SkipListSet[K constraints.Ordered] interface {
	Insert(key K)

	Contains(key K) bool

	Delete(key K)

	Size() uint
}

func NewSkipListSet[K constraints.Ordered]() SkipListSet[K] {
	return newSkipListSet[K]()
}

func newSkipListSet[K constraints.Ordered]() *skipListSet[K] {
	head := node.NewSentinelNode[K](MaxHeight, MaxHeight)
	head.Height = MaxHeight
	head.IsHead = true

	tail := node.NewSentinelNode[K](MaxHeight, MaxHeight)
	tail.Height = MaxHeight
	tail.IsTail = true
	for i := 0; i < MaxHeight; i++ {
		head.Next[i] = tail
		tail.Next[i] = nil
	}
	return &skipListSet[K]{
		head: head,
		tail: tail,
		size: 0,
		rng:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
