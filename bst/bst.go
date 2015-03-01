package bst

// BST a binary search tree which uses strings as keys
type BST struct {
	left, right *BST
	node        interface{}
	Compare     CompareFunc
}

func NewBST(val interface{}, comp CompareFunc) *BST {
	return &BST{
		node:    val,
		Compare: comp,
	}
}

// CompareFunc compare two interfaces
// i1 < i2 = -1
// i1 > i2 = 1
// i1 == i2 = 0
type CompareFunc func(i1, i2 interface{}) int

// Search find a value in the tree given a key
func (b *BST) Search(i interface{}) interface{} {
	dir := b.Compare(b.node, i)

	switch dir {
	case 0:
		return b.node
	case -1:
		if b.right == nil {
			return nil
		}
		return b.right.Search(i)
	case 1:
		if b.left == nil {
			return nil
		}
		return b.left.Search(i)
	default:
		panic("invalid compare function")
	}
}

// Exists check for a keys existance
func (b *BST) Exists(key interface{}) bool {
	return b.Search(key) != nil
}

// Insert add a key,val pair into the BST
func (b *BST) Insert(val interface{}) {
	dir := b.Compare(b.node, val)

	switch dir {
	case 0:
		b.node = val
	case 1:
		if b.left == nil {
			b.left = NewBST(val, b.Compare)
		} else {
			b.left.Insert(val)
		}

	case -1:
		if b.right == nil {
			b.right = NewBST(val, b.Compare)
		} else {
			b.right.Insert(val)
		}
	default:
		panic("invalid compare function")
	}
}
