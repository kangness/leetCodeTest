package main

import (
	"container/list"
	"math"
)

type MinStack struct {
	min *list.Element
	data *list.List
}


/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{min: nil}
	m.data = list.New()
	m.data.Init()
	return m
}


func (this *MinStack) Push(x int)  {
	this.data.PushFront(x)
	f := this.data.Front()
	if this.min == nil {
		this.min = f
		return
	}
	if this.min.Value.(int) > x {
		this.min = f
	}
	return
}


func (this *MinStack) Pop()  {
	f := this.data.Front()
	if f != nil {
		this.data.Remove(f)
	}
	if f != nil && this.min == f {
		this.min = this.data.Front()
		for e := this.data.Front(); e != nil; e = e.Next() {
			if this.min.Value.(int) > e.Value.(int) {
				this.min = e
			}
		}
	}
}


func (this *MinStack) Top() int {
	if this.data != nil {
		f := this.data.Front()
		if f != nil {
			return f.Value.(int)
		}
	}
	return 0
}


func (this *MinStack) GetMin() int {
	if this.min != nil {
		return this.min.Value.(int)
	}
	return 0
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
