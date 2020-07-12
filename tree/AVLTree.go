package tree
import (
	"fmt"
	"reflect"
	//"math/rand"
	//"time"
)

type AVLTree struct {
	BaseTreeManager
}

func (self *AVLTree) createNode(value interface{}, left, right *BaseTree) (*BaseTree) {
	return &BaseTree{
		Value: value,
		Left: left,
		Right: right,
		Hight: 0,
	}
}

func (self *AVLTree) getHight(node *BaseTree) int {
	if node == nil {
		return 0
	}
	return node.Hight
}

func (self *AVLTree) Max(a, b interface{}) bool {
	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return true
		}
		return false
	case string:
		if a.(string) > b.(string) {
			return true
		}
		return false
	default:
		fmt.Sprintf("not support type ", reflect.TypeOf(a))
	}
	return false
}
func (self *AVLTree) Enqual(a, b interface{}) bool {
	switch a.(type) {
	case int:
		return  a.(int) == b.(int)
	case string:
		return a.(string) == b.(string)
	default:
		fmt.Sprintf("not support type ", reflect.TypeOf(a))
	}
	return false
}
func (self *AVLTree) setNodeHight(node *BaseTree) {
	if node == nil {
		return
	}
	if self.getHight(node.Left) > self.getHight(node.Right) {
		node.Hight = self.getHight(node.Left) + 1
	}else {
		node.Hight = self.getHight(node.Right) + 1
	}
}
func (self *AVLTree) LLRotation(node *BaseTree) *BaseTree{
	var b *BaseTree
	b = node.Left
	node.Left = b.Right
	b.Right = node
	self.setNodeHight(node)
	self.setNodeHight(b)
	return b
}

func (self *AVLTree) RRRotation(node *BaseTree) *BaseTree {
	var c *BaseTree
	c = node.Right
	node.Right = c.Left
	c.Left = node
	self.setNodeHight(node)
	self.setNodeHight(c)
	return c
}

func (self *AVLTree) LRRotation(node *BaseTree) *BaseTree {
	node.Left = self.RRRotation(node.Left)
	return self.LLRotation(node)
}

func (self *AVLTree) RLRotation(node *BaseTree) *BaseTree {
	node.Right = self.LLRotation(node.Right)
	return self.RRRotation(node)
}

func (self *AVLTree) getNodeHightOffset(n1, n2 *BaseTree) int {
	var a, b int
	if n1 != nil  {
		a = n1.Hight
	}
	if n2 != nil {
		b = n2.Hight
	}
	return a - b
}
func (self *AVLTree) InsertNode(node *BaseTree, value interface{}) *BaseTree {
	return self.insertNode(node, value)
}

func (self *AVLTree) deleteNode(node *BaseTree, value interface{}) *BaseTree {
	if node == nil {
		return node
	}
	if self.Enqual(node.Value, value) {
		if node.Left != nil && node.Right != nil {

		}
	}
}
func (self *AVLTree) insertNode(node *BaseTree, value interface{}) *BaseTree {
	if node == nil {
		return self.createNode(value,nil,nil)
	}
	if self.Max(node.Value, value) {
		node.Left = self.insertNode(node.Left,value)
		offset := self.getNodeHightOffset(node.Right, node.Left)
		if offset >= 2 || offset <= -2 {
			if self.Max(node.Left.Value,value) {
				node = self.LLRotation(node)
			}else {
				node = self.LRRotation(node)
			}
		}
	}else if self.Max(value, node.Value){
		node.Right = self.insertNode(node.Right, value)
		offset := self.getNodeHightOffset(node.Left, node.Right)
		if offset >= 2 || offset <= -2 {
			if self.Max(value,node.Right.Value) {
				node = self.RRRotation(node)
			}else {
				node = self.RLRotation(node)
			}
		}
	}else {
		fmt.Println("node value is the same ", value)
	}
	self.setNodeHight(node)
	return node
}

func (self *AVLTree) Find(node *BaseTree, value interface{}) *BaseTree {
	for {
		if node == nil {
			return nil
		}
		if self.Enqual(node.Value, value) {
			return node
		}
		if self.Max(node.Value, value) {
			node = node.Left
		}else {
			node = node.Right
		}
	}
	return nil
}
/*
func main() {
	manager := AVLTree{}
	rand.Seed(time.Now().UnixNano())
	var node *BaseTree
	for i:= 0; i <= 20; i ++ {
		num := rand.Intn(1000)
		node = manager.insertNode(node,num)
	}
	fmt.Println(manager.LDR(node) )
}

 */
