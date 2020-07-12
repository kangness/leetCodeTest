package main
import (
	"container/list"
	"fmt"
)
type LRUCache struct {
   maxEntry int
   ll *list.List
   data map[int] _entry
    
}
type _entry struct {
	key int
	value int
	el *list.Element
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
    	maxEntry: capacity,
    	ll: list.New(),
    	data: make(map[int]_entry),
    }
}

func (this *LRUCache) Get(key int) int {
    value, ok := this.data[key]
    if !ok {
    	return -1
    }
    this.ll.MoveToFront(value.el)
    return value.el.Value.(_entry).value
}


func (this *LRUCache) Put(key int, value int)  {
	v, ok := this.data[key]
	if !ok {
		if this.ll.Len() >= this.maxEntry {
			fmt.Println(key, value)
			back := this.ll.Back()
			delete(this.data,back.Value.(_entry).key)
			fmt.Println(this.data, this.ll.Len())
			this.ll.Remove(back)
		}
	} else {
		this.ll.Remove(v.el)
	}
	newEntry := _entry{
			key : key,
			value:value,
			el: nil,
	}
	newEntry.el = this.ll.PushFront(newEntry)
	this.data[key] = newEntry
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1,1)
	obj.Put(2,2,)
	obj.Get(1)
	obj.Put(3,3)
	obj.Get(2)
}