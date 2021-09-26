package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	iter := l
	for iter != nil {
		if count == pos {
			return iter
		}
		iter = iter.Next
		count++
	}
	return nil
}
