package list

import (
	"fmt"
	"strings"
)

func (s *skipList) String() string {
	var sb strings.Builder
	for i := MaxHeight - 1; i >= 0; i -= 1 {
		sb.WriteString("-\u221e ")
		curr := s.head.Next[i]
		for j := 0; j < int(s.size); j++ {
			if curr != nil && curr != s.tail {
				sb.WriteString(fmt.Sprintf("- > %d -", curr.Val))
			} else {
				sb.WriteString(" -")
			}
			curr = curr.Next[i]
		}
		sb.WriteString(" - > \u221e\n")
	}
	return sb.String()
}
