package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruItem struct {
	key   Key
	value any
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value any) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if i, ok := l.items[key]; ok {
		i.Value = value
		l.queue.MoveToFront(i)
		l.items[key].Value = lruItem{key: key, value: value}
		return true
	}

	l.removeOld()
	l.items[key] = l.queue.PushFront(lruItem{key: key, value: value})
	return false
}

func (l *lruCache) Get(key Key) (any, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if i, ok := l.items[key]; ok {
		l.queue.MoveToFront(i)
		return i.Value.(lruItem).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func (l *lruCache) full() bool {
	return len(l.items) >= l.capacity
}

func (l *lruCache) removeOld() {
	if l.full() {
		e := l.queue.Back()
		delete(l.items, e.Value.(lruItem).key)
		l.queue.Remove(e)
	}
}
