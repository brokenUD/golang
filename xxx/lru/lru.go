package lru

import ("container/list"
"log")

type Cache struct{
	max int64
	usered int64
	ll *list.List
	cache map[string]*list.Element
	onEvict func(key string, value Value)
}

type entry struct{
	key string
	value Value
}

type Value interface {
	Len() int
}

func New(maxByte int64, onEvict func(string, Value))*Cache{
	return &Cache{
		max: maxByte,
		ll: list.New(),
		cache: make(map[string]*list.Element),
		onEvict: onEvict,
	}
}

func (c *Cache)Get(key string)(value Value, ok bool){
	if ele, ok := c.cache[key];ok{
		c.ll.MoveToBack(ele)
		kv:=ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache)RemoveOld() {
	ele := c.ll.Front()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache,kv.key)
		c.usered -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.onEvict!=nil {
			c.onEvict(kv.key,kv.value)
		}
	}
}

func (c *Cache)Add(key string, value Value){
	if int64(value.Len())>c.max{
		log.Printf("over max len")
		return
	}
	if ele,ok := c.cache[key];ok {
		c.ll.MoveToBack(ele)
		kv := ele.Value.(*entry)
		c.usered += int64(value.Len()) - int64(kv.value.Len())
	} else {
		ele := c.ll.PushFront(&entry{key: key,value: value})
		c.cache[key] = ele
		c.usered += int64(len(key)) + int64(value.Len())
	}
	for c.max<c.usered {
		c.RemoveOld()
	}
	return
}
