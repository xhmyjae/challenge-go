package piscine

func ListLast(l *List) interface{} {
	b := l.Head
	if b == nil {
		return nil
	}
	for b.Next != nil {
		b = b.Next
	}
	return b.Data
}
