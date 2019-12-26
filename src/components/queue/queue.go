package queue

import (
	"container/list"
	"sync"
)

type Queue struct {
	list  *list.List
	Mutex sync.RWMutex
}

func NewQueue() *Queue {
	list := list.New()
	return &Queue{list, sync.RWMutex{}}
}

func (queue *Queue) Add(value interface{}) {
	queue.Mutex.Lock()
	queue.list.PushBack(value)
	queue.Mutex.Unlock()
}

func (queue *Queue) Remove() interface{} {
	queue.Mutex.Lock()
	defer queue.Mutex.Unlock()

	e := queue.list.Front()
	if e != nil {
		queue.list.Remove(e)
		return e.Value
	}
	return nil
}

func (queue *Queue) Peek() interface{} {
	queue.Mutex.RLock()
	defer queue.Mutex.RUnlock()

	e := queue.list.Front()
	if e != nil {
		return e.Value
	}
	return nil
}

func (queue *Queue) Length() int {
	queue.Mutex.RLock()
	defer queue.Mutex.RUnlock()

	return queue.list.Len()
}

func (queue *Queue) Empty() bool {
	queue.Mutex.RLock()
	defer queue.Mutex.RUnlock()

	return queue.list.Len() == 0
}
