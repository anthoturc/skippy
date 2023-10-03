package node

import "golang.org/x/exp/constraints"

type SkipNode[K constraints.Ordered] struct {
	Key    K
	Next   []*SkipNode[K]
	Height int
	IsHead bool
	IsTail bool
}

func NewSkipNode[K constraints.Ordered](val K, height, maxHeight int) *SkipNode[K] {
	next := make([]*SkipNode[K], maxHeight)
	return &SkipNode[K]{
		Key:    val,
		Next:   next,
		Height: height,
		IsHead: false,
		IsTail: false,
	}
}

func NewSentinelNode[K constraints.Ordered](height, maxHeight int) *SkipNode[K] {
	next := make([]*SkipNode[K], maxHeight)
	var sentinelVal K
	return &SkipNode[K]{
		Key:    sentinelVal,
		Next:   next,
		Height: height,
	}
}
