package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	a := &NodeL{Data: data} // créer le node qu'on veut à la fin et le & car list a un type pointeur
	b := l.Head
	if b == nil {
		l.Head = a
		return
	}
	for b.Next != nil {
		b = b.Next
	}
	b.Next = a
}
