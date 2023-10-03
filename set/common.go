package set

import (
	"github.com/anthoturc/skippy/internal/node"
)

func (s *skipListSet[K]) searchKey(key K) *node.SkipNode[K] {
	curr := s.head
	idx := s.head.Height - 1
	for idx >= 0 && curr != nil {
		for idx > 0 && s.lessEq(key, curr.Next[idx]) {
			idx -= 1
		}

		// If this was the case then we would have kept going down in the loop
		// but we are on the bottom row instead and there is no other place to go.
		// In this case we have found the predecessor or the correct node to insert
		if s.lessEq(key, curr.Next[idx]) {
			break
		}
		curr = curr.Next[idx]
	}
	return curr.Next[idx]
}

func (s *skipListSet[K]) searchPredecessorsForKey(key K) []*node.SkipNode[K] {

	curr := s.head
	idx := s.head.Height - 1

	predecessors := make([]*node.SkipNode[K], s.head.Height)
	for idx >= 0 && curr != nil {
		// This condition ensures that we go through the entire height of the list.
		// If the key was strictly less than the next value at a given height (i.e. the idx),
		// then we would stop early and some members of the predecessor would be nil.
		for idx > 0 && s.lessEq(key, curr.Next[idx]) {
			predecessors[idx] = curr
			idx -= 1
		}

		// If this was the case then we would have kept going down in the loop
		// but we are on the bottom row instead and there is no other place to go.
		// In this case we have found the predecessor or the correct node to insert if the current
		// node's *next* value is greater than or equal to the key
		if s.lessEq(key, curr.Next[idx]) {
			break
		}
		curr = curr.Next[idx]
	}
	// After the loop, idx should always be zero
	predecessors[idx] = curr

	return predecessors
}

func (s *skipListSet[K]) pick50Fifty() int {
	return s.rng.Intn(100)
}

func (s *skipListSet[K]) genRandomHeight() (height int) {
	height = 1
	for height < MaxHeight && s.pick50Fifty() < 65 { // TODO: The probability should be put into a skip list set config.
		height += 1
	}
	return height
}

// lessEq will return true if the `key` <= `n.Val`. This function accounts
// for the head and tail being possible comparisons
func (s *skipListSet[K]) lessEq(key K, n *node.SkipNode[K]) bool {
	if n.IsHead {
		// The head corresponds to -inf for the underlying type
		// so it will *always* be less than
		return false
	}

	if n.IsTail { // The tail corresponds to +inf for the underlying type
		return true
	}

	return key <= n.Key
}
