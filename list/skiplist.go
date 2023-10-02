package list

import (
	"github.com/anthoturc/go-skippy/internal/node"
)

// Insert will attempt to add an item to the SkipList if it doesn't already exist.
func (s *skipList) Insert(item int) {

	predecessors := s.searchPredecessorsForInsert(item)
	if predecessors[0].Val == item {
		return // The item is already in the list so
	}

	height := s.genRandomHeight()
	additionalNode := node.NewSkipNode(item, height, MaxHeight)
	for i := 0; i < height; i++ {
		additionalNode.Next[i] = predecessors[i].Next[i]
		predecessors[i].Next[i] = additionalNode
	}

	s.size += 1
}

func (s *skipList) Find(item int) int {
	node := s.searchKey(item)
	if node != nil && node != s.head && node != s.tail {
		return node.Val
	}
	return 0
}

func (s *skipList) Contains(item int) bool {
	node := s.searchKey(item)

	return node != nil &&
		// There could be a case where the item's value is MinInt but we don't want to confuse that
		// with the tail
		node != s.tail &&
		node.Val == item
}

func (s *skipList) Delete(item int) {

}

func (s *skipList) Size() uint {
	return s.size
}
