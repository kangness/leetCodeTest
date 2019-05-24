package main

import "fmt"

type BaseTree struct {
	Value interface{}
	Left  *BaseTree
	Right *BaseTree
}


type BaseTreeManager struct {

}
// 前序遍历
func (self *BaseTreeManager) DLR(node *BaseTree) []interface{} {
	var result []interface{}
	if node == nil {
		return result
	}
	head := node
	if head.Value != nil {
		result = append(result, head.Value)
	}
	if head.Left != nil {
		result = append(result, self.DLR(head.Left)...)
	}
	if head.Right != nil {
		result = append(result, self.DLR(head.Right)...)
	}
	return result
}
// 中序遍历
func (self *BaseTreeManager) LDR(node *BaseTree) []interface{} {
	var result []interface{}
	if node == nil {
		return result
	}
	if node.Left != nil {
		result = append(result, self.LDR(node.Left)...)
	}
	if node.Value != nil {
		result = append(result, node.Value)
	}
	if node.Right != nil {
		result = append(result, self.LDR(node.Right)...)
	}
	return result
}
func (self *BaseTreeManager) LRD (node *BaseTree) []interface{} {
	var result []interface{}
	if node == nil {
		return result
	}
	if node.Left != nil {
		result = append(result, self.LRD(node.Left)...)
	}
	if node.Right != nil {
		result = append(result, self.LRD(node.Right)...)
	}
	if node.Value != nil {
		result = append(result, node.Value)
	}
	return result
}
func main() {
	manager := &BaseTreeManager{}
	node := &BaseTree{
		Value: interface{}("A"),
	}
	b  := &BaseTree{
		Value: interface{}("B"),
	}
	e := &BaseTree{
		Value: interface{}("E"),
	}
	c := &BaseTree{
		Value:interface{}("C"),
	}
	f := &BaseTree{
		Value: interface{}("F"),
	}
	d := &BaseTree{
		Value: interface{}("D"),
	}
	g := &BaseTree{
		Value: interface{}("G"),
	}
	h := &BaseTree{
		Value: interface{}("H"),
	}
	k := &BaseTree{
		Value: interface{}("K"),
	}
	node.Left = b
	node.Right = e
	b.Right = c
	e.Right = f
	c.Left = d
	f.Left = g
	g.Left = h
	g.Right = k
	fmt.Println(manager.DLR(node))
	fmt.Println(manager.LDR(node))
	fmt.Println(manager.LRD(node))
}
