package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	iter := l.Head
	for iter != nil {
		if cond(iter) {
			f(iter)
		}
		iter = iter.Next
	}

}
