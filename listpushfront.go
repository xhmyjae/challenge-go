package piscine

func ListPushFront(l *List, data interface{}) {
	a := &NodeL{Data: data}
	b := l.Head
	if b == nil {
		l.Head = a
		return
	}
	a.Next = b
	l.Head = a
}
