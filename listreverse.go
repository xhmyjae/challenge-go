package piscine

func ListReverse(l *List) {
	n := &List{}
	iter := l.Head
	for iter != nil {
		ListPushFront(n, iter.Data)
		iter = iter.Next
	}
	*l = *n
}
