package tree

import "testing"

func intCompare(i1 interface{}, i2 interface{}) int {
	f := i1.(*testContainer)
	s := i2.(*testContainer)

	if f.key == s.key {
		return 0
	}

	if f.key < s.key {
		return -1
	}

	return 1
}

type testContainer struct {
	key int
	val interface{}
}

func TestInsertAndSearch(t *testing.T) {
	bst := NewBST(&testContainer{
		key: 100,
		val: "hello",
	}, intCompare)

	bst.Insert(&testContainer{
		key: 200,
		val: "world",
	})

	re := bst.Search(&testContainer{
		key: 200,
	})

	if re.(*testContainer).val != "world" {
		t.Fail()
	}
}

func TestExists(t *testing.T) {
	bst := NewBST(&testContainer{
		key: 1000,
		val: nil,
	}, intCompare)

	if bst.Exists(&testContainer{key: 1}) {
		t.Fail()
	}

	if !bst.Exists(&testContainer{key: 1000}) {
		t.Fail()
	}
}
