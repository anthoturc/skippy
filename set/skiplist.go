package set

import (
	"github.com/anthoturc/skippy/internal/node"
)

// Insert will attempt to add an item to the SkipList if it doesn't already exist.
func (s *skipListSet) Insert(key int) {

	predecessors := s.searchPredecessorsForKey(key)

	// There is a chance that the first predecessor's
	// immediate successor *is* the item we are trying to insert
	if predecessors[0].Next[0].Val == key {
		return
	}

	height := s.genRandomHeight()
	additionalNode := node.NewSkipNode(key, height, MaxHeight)
	for i := 0; i < height; i++ {
		additionalNode.Next[i] = predecessors[i].Next[i]
		predecessors[i].Next[i] = additionalNode
	}

	s.size += 1
}

func (s *skipListSet) Contains(key int) bool {
	// Don't bother searching for the key if there aren't any
	if s.size == 0 {
		return false
	}
	node := s.searchKey(key)

	return node != nil &&
		// There could be a case where the item's value is MinInt but we don't want to confuse that
		// with the tail
		node != s.tail &&
		node.Val == key
}

func (s *skipListSet) Delete(key int) {
	// Don't bother looking for the key if there aren't any
	if s.size == 0 {
		return
	}

	predecessors := s.searchPredecessorsForKey(key)

	// If the key doesn't exist at the lowest level of the skip list -- i.e. the successor of the predcessor
	// at level 0, then the key doesn't exist at all so the delete is a no-op.
	if predecessors[0].Next[0].Val != key {
		return
	}

	keyNode := predecessors[0].Next[0]
	keyHeight := keyNode.Height
	for i := 0; i < keyHeight; i++ {
		predecessors[i].Next[i] = keyNode.Next[i]
	}
	s.size -= 1
}

func (s *skipListSet) Size() uint {
	return s.size
}
