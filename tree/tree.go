package tree

// Tree
type Tree interface {
	Insert(interface{})
	Exists(interface{}) bool
	Empty(interface{}) bool
	Search(interface{}) interface{}
}
