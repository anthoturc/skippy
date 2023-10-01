package node

type SkipNode struct {
	Val    int
	Next   []*SkipNode
	Height int
}

func NewSkipNode(val, height, maxHeight int) *SkipNode {
	next := make([]*SkipNode, maxHeight)
	return &SkipNode{
		Val:    val,
		Next:   next,
		Height: height,
	}
}
