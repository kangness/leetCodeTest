package tree
import (
	"container/list"
)

type Cache struct {
	MaxEntries int
	ll *list.List
	cache map[string]*list.Element
}

type entry struct {
	key string
	value interface{}
}

func New(max int) *Cache {
	return &Cache{
		MaxEntries: max,
		ll: list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (self *Cache) Set(key string, value interface{}) bool {
	if self.cache == nil {
		self.cache = make(map[string]*list.Element)
	}
	e, ok := self.cache[key]
	if ok {
		self.ll.MoveToFront(e)
		e.Value.(*entry).value = value
		return true
	}
	ele := self.ll.PushFront(&entry{
		key: key,
		value: value,
	})
	self.cache[key] = ele
	if self.MaxEntries != 0 && self.MaxEntries < self.ll.Len() {
		self.removeOldest()
	}
	return true
}

func (self *Cache) Get(key string) (interface{},bool) {
	if self.cache == nil {
		return nil, false
	}
	e, ok := self.cache[key]
	if !ok {
		return nil, false
	}
	self.ll.MoveToFront(e)
	return e.Value.(*entry).value,true
}
func (self *Cache) removeOldest() {
	if self.cache == nil {
		return
	}
	self.removeElement(self.ll.Back())
}

func (self *Cache) Remove(key string) {
	if self.cache == nil {
		return
	}
	e, ok := self.cache[key]
	if ok {
		self.removeElement(e)
	}
}
func (self *Cache) removeElement(value *list.Element) {
	if value == nil {
		return
	}
	key := value.Value.(*entry).key
	delete(self.cache, key)
	self.ll.Remove(value)
}

