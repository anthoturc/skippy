package list

import "github.com/anthoturc/go-skippy/internal/node"

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
		// In this case we have found the predecessor or the correct node to insert
		if key <= curr.Val {
			break
		}
		curr = curr.Next[idx]
	}
	// After the loop, idx should always be zero
	predecessors[idx] = curr

	return predecessors
}
