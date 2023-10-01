package list

import (
	"math/rand"
	"time"

	"github.com/anthoturc/go-skippy/internal/node"
)

func (s *skipList) searchKey(key int) *node.SkipNode {
	node := s.head
	idx := s.head.Height - 1
	for idx >= 0 && node != nil {
		for idx > 0 && key < node.Next[idx].Val {
			idx -= 1
		}

		// If this was the case then we would have kept going down in the loop
		// but we are on the bottom row instead and there is no other place to go.
		// In this case we have found the predecessor or the correct node to insert
		if key <= node.Next[idx].Val {
			break
		}
		node = node.Next[idx]
	}
	return node
}

func (s *skipList) searchPredecessorsForInsert(key int) []*node.SkipNode {

	curr := s.head
	idx := s.head.Height - 1

	predecessors := make([]*node.SkipNode, s.head.Height)
	for idx >= 0 && curr != nil {
		for idx > 0 && key < curr.Next[idx].Val {
			predecessors[idx] = curr
			idx -= 1
		}

		// If this was the case then we would have kept going down in the loop
		// but we are on the bottom row instead and there is no other place to go.
		// In this case we have found the predecessor or the correct node to insert if the current
		// node's *next* value is greater than or equal to the key
		if key <= curr.Next[idx].Val {
			break
		}
		curr = curr.Next[idx]
	}
	// After the loop, idx should always be zero
	predecessors[idx] = curr

	return predecessors
}

func (s *skipList) pick50Fifty() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
}

func (s *skipList) genRandomHeight() (height int) {
	height = 1
	for height < MaxHeight && s.pick50Fifty() < 50 {
		height += 1
	}
	return height
}
