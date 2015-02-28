package list

// LinkedList an implementation of a doubly linked list
type LinkedList struct {
	first  *node
	last   *node
	length int
}

// Push push an element on the front of the list
func (l *LinkedList) Push(i interface{}) {
	n := &node{
		elem: i,
	}
	if l.first != nil {
		l.first.addBefore(n)
	}
	l.first = n
	l.length++
}

// Insert add an element to the list at a given index
func (l *LinkedList) Insert(elem interface{}, index int) {
	if index > l.length {
		return
	}

	// seek to the element we want
	n := l.first
	for i := 0; i < index; i++ {
		n = n.next
	}
	n.addBefore(&node{
		elem: elem,
	})
}

// Index get an element at a given index
func (l *LinkedList) Index(index int) interface{} {
	if index > l.length {
		return nil
	}

	// seek to "index" in the list
	n := l.first
	for i := 0; i < index; i++ {
		n = n.next
	}

	return n.elem
}

// Pop pop an elemnt off the first of the list
func (l *LinkedList) Pop() interface{} {
	tmp := l.first
	l.first = tmp.next

	if l.first != nil {
		l.first.prev = nil
	}
	return tmp.elem
}

// Peek get the first element of the list without removing it
func (l *LinkedList) Peek() interface{} {
	if l.first == nil {
		return nil
	}
	return l.first.elem
}

// Delete remove a given element from the list
func (l *LinkedList) Delete(i interface{}) {
	f := l.first

	// for every element in the list, if the element given is found, remove it from the list
	for f != nil {
		if f.elem == i {
			if f.prev != nil {
				f.prev.next = f.next
			}

			if f.next != nil {
				f.next.prev = f.prev
			}

			if f.first() {
				l.first = f.next
			}

			if f.last() {
				l.last = f.prev
			}
			// we now have one less element
			l.length -= 1
		}
		f = f.next
	}
}

type node struct {
	prev, next *node
	elem       interface{}
}

func (n *node) first() bool {
	return n.prev == nil
}

func (n *node) last() bool {
	return n.next == nil
}

func (n *node) addBefore(e *node) {
	if n.prev != nil {
		n.prev.next = e
	}
	e.prev = n.prev
	n.prev = e
	e.next = n
}

func (n *node) addAfter(e *node) {
	if n.next != nil {
		n.next.prev = e
	}
	e.next = n.next
	n.next = e
	e.prev = n
}
