package set

import "testing"

func TestAdd(t *testing.T) {
	s := New()
	s.Add("hello")
	if !s.Test("hello") {
		t.Fail()
	}

	if s.Test("not-there") {
		t.Fail()
	}
}
