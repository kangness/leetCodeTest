package main

import "container/list"

type MyStack struct {
	myList *list.List
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	s := MyStack{}
	s.myList = list.New()
	return s
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.myList.PushFront(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	f := this.myList.Front()
	if f == nil {
		return 0
	}
	value := f.Value.(int)
	this.myList.Remove(f)
	return value
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.myList.Front().Value.(int)
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.myList == nil || this.myList.Len() <= 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
