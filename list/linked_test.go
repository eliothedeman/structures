package list

import "testing"

func TestPush(t *testing.T) {
	l := &LinkedList{}
	for i := 0; i < 20; i++ {
		l.Push(i)
	}

	if l.Peek() != 19 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	l := &LinkedList{}
	for i := 0; i < 20; i++ {
		l.Push(i)
	}

	for i := 0; i < 20; i++ {
		l.Pop()
	}

	if l.Peek() != nil {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	l := &LinkedList{}
	for i := 0; i < 100; i++ {
		l.Push(i)
	}
	l.Insert(999, 4)

	if l.Index(4) != 999 {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	l := &LinkedList{}
	for i := 213; i >= 0; i-- {
		l.Push(i)
	}

	org := l.Index(21)

	l.Delete(21)

	if org == l.Index(21) {
		t.Fail()
	}
}

func TestWalk(t *testing.T) {
	l := &LinkedList{}

	for i := 0; i < 10; i++ {
		l.Push(i)
	}

	l.Reset()
	for i := 0; i < 10; i++ {
		if l.Next() != 9-i {
			t.Fail()
		}
	}
}
