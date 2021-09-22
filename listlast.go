package piscine

func ListLast(l *List) interface{} {
	b := l.Head
	for b != nil {
		b = b.Next
	}
	return b
}
