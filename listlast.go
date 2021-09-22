package piscine

func ListLast(l *List) interface{} {
	b := l.Head
	for b.Next != nil {
		b = b.Next
	}
	return b.Data
}
