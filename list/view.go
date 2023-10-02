package list

import (
	"fmt"
	"strings"
)

func (s *skipList) String() string {

	index := map[int]int{}
	node := s.head.Next[0]
	for i := 0; i < int(s.size); i++ {
		index[i] = node.Val
		node = node.Next[0]
	}

	var sb strings.Builder
	for i := MaxHeight - 1; i >= 0; i -= 1 {
		sb.WriteString("-\u221e -")
		curr := s.head.Next[i]
		for j := 0; j < int(s.size); j++ {
			if curr != nil && curr.Val == index[j] {
				sb.WriteString(fmt.Sprintf(" %d ", curr.Val))
				curr = curr.Next[i]
			} else {
				sb.WriteString(" - ")
			}
		}
		sb.WriteString("- > \u221e\n")
	}
	return sb.String()
}
