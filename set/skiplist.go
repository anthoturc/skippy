package set

import (
	"github.com/anthoturc/skippy/internal/node"
)

// Insert will attempt to add an item to the SkipList if it doesn't already exist.
func (s *skipListSet[K]) Insert(key K) {

	predecessors := s.searchPredecessorsForKey(key)

	// There is a chance that the first predecessor's
	// immediate successor *is* the item we are trying to insert
	// Note: The zero value for a given type should still be inserted
	// even though, the value corresponds to the head and tail of the list
	if predecessors[0].Next[0] != s.tail && predecessors[0].Next[0].Key == key {
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

func (s *skipListSet[K]) Contains(key K) bool {
	// Don't bother searching for the key if there aren't any
	if s.size == 0 {
		return false
	}
	node := s.searchKey(key)
	return node != nil &&
		node != s.tail &&
		node.Key == key
}

func (s *skipListSet[K]) Delete(key K) {
	// Don't bother looking for the key if there aren't any
	if s.size == 0 {
		return
	}

	predecessors := s.searchPredecessorsForKey(key)

	// If the key doesn't exist at the lowest level of the skip list -- i.e. the successor of the predcessor
	// at level 0, then the key doesn't exist at all so the delete is a no-op.
	if predecessors[0].Next[0].Key != key {
		return
	}

	keyNode := predecessors[0].Next[0]
	keyHeight := keyNode.Height
	for i := 0; i < keyHeight; i++ {
		predecessors[i].Next[i] = keyNode.Next[i]
	}
	s.size -= 1
}

func (s *skipListSet[K]) Size() uint {
	return s.size
}
