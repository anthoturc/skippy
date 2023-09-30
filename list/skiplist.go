package list

func (s *skipList) Insert(item int) {

}

func (s *skipList) Find(item int) int {
	return 0
}

func (s *skipList) Contains(item int) bool {
	node := s.searchKey(item)

	return node != nil &&
		node != s.head && // We don't want to count the head or the tail even if asked for MaxInt or MinInt
		node != s.tail &&
		node.Val == item
}

func (s *skipList) Delete(item int) {

}
