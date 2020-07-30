package tree

import "github.com/google/btree"

const Red int = 1
const Black int = 2
type Node struct {
	Parent *Node
	Left *Node
	Right *Node
	Color int
	Item interface{}
}

type RBTTree struct {
	NIL *Node
	Cur *Node
	Root *Node
	Count int
}

func New() (*RBTTree) {
	node := &Node{nil, nil, nil, Black, nil}
	return &RBTTree{
		NIL: node,
		Cur: nil,
		Root: node,
		Count: 0,
	}
}

func (rbt *RBTTree) LeftRotate(no *Node) {

}

func (rbt *RBTTree) RightRotate(node *Node) {

}

func (rbt *RBTTree) Insert(node *Node) {

}

func (rbt *RBTTree) Delete(node *Node) {

}

func (rbt *RBTTree) Search (value interface{}) *Node {

}
