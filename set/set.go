package set

import (
	"github.com/eliothedeman/structures/tree"
)

// An implementation of set for strings
type Set struct {
	tree *tree.BST
}

// Add add a string to the set
func (s *Set) Add(key string) {
	s.tree.Insert(key)
}

// Test test for set membership
func (s *Set) Test(key string) bool {
	return s.tree.Exists(key)
}

func New() *Set {
	return &Set{
		tree: tree.NewBST(nil, compareString),
	}
}

func compareString(i1, i2 interface{}) int {
	s1 := i1.(string)
	s2 := i2.(string)
	// determin the shortest string
	var l int
	l1 := len(s1)
	l2 := len(s2)
	if l1 < l2 {
		l = l1
	} else {
		l = l2
	}

	for i := 0; i < l; i++ {
		if s1[i] > s2[i] {
			return 1
		}
		if s1[i] < s2[i] {
			return -1
		}
	}

	if l1 == l2 {
		return 0
	}

	return -1
}
