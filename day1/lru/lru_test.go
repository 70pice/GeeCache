package lru

import (
	"fmt"
	"testing"
)

type String string

func (d String) len() int {
	return len(d)
}

func TestAdd(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("qwe"))
	lru.Add("key2", String("asd"))
	lru.Add("key3", String("zxc"))
	for item := lru.ll.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("qwe"))
	lru.Add("key2", String("asd"))
	lru.Add("key3", String("zxc"))
	lru.RemoveOldest()
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("qwe"))
	lru.Add("key2", String("asd"))
	lru.Add("key3", String("zxc"))

	lru.Get("key1")
}