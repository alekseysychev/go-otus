package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	sync.Mutex
	queue List
	items map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	var item *ListItem
	var ok bool
	var finded bool = false

	c.Mutex.Lock()
	if item, ok = c.items[key]; !ok { // not finded, need insert
		if c.queue.Len() >= c.capacity {
			lastItem := c.queue.Back()
			c.queue.Remove(lastItem)

			for lastItemKey, lastItemValue := range c.items {
				if lastItemValue == lastItem {
					delete(c.items, lastItemKey)
					break
				}
			}
		}

		item = c.queue.PushFront(value)
	} else {
		item.Value = value
		c.queue.MoveToFront(item)
		finded = true
	}
	c.items[key] = item
	c.Mutex.Unlock()
	return finded
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	var item *ListItem
	var ok bool
	var value interface{}
	c.Mutex.Lock()
	if item, ok = c.items[key]; ok {
		c.queue.MoveToFront(item)
		value = item.Value
	}
	c.Mutex.Unlock()

	return value, ok
}

func (c *lruCache) Clear() {
	c.Mutex.Lock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.Mutex.Unlock()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
