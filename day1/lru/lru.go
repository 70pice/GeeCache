package lru

import (
	"container/list"
)

type Cache struct {
	// maxBytes允许使用的最大内存
	maxBytes int64
	// nbytes是当前已经使用的内存
	nbytes int64
	// list.list是Go标准库实现的双向链表
	ll *list.List
	// 键是字符串，值是双向链表中对应的指针
	cache map[string]*list.Element
	// 某条记录被移除时的回调函数，我们允许所有类型的数据，只要这个数据返回所占用的内存大小
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找的功能优 1. 从字典中找到对应的双向链表节点 2. 把所访问的节点移动到队尾（移除的时候就可以移除对头节点）

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 这里的删除实际上是缓存淘汰，移除最少访问的节点，队首
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		// 原有的键是否存在，存在进入这个逻辑
		// 如果存在将元素移动到队首
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		// nbytes是已经使用的内存大小，因为原有的value可能会被替换，内存重新计算
		c.nbytes += int64(value.len()) - int64(kv.value.len())
		kv.value = value
	} else {
		// 创建一个新的接地那插入，在队首插入
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		// 插入之后需要更新内存，如果内存超出限制，应该削减内存
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
