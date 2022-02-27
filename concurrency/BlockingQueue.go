package concurrency

import (
	"sync"
	"sync/atomic"
)

type BlockingQueue struct {
	head, tail     int
	capacity, size uint64
	data           []int
	mu             sync.Mutex
	enq, dq        chan int
}

func NewBlockingQueue(capacity uint64) *BlockingQueue {
	return &BlockingQueue{
		head:     -1,
		tail:     -1,
		capacity: capacity,
		data:     make([]int, capacity),
		size:     0,
		enq:      make(chan int, 1),
		dq:       make(chan int, 1),
	}
}

func (b *BlockingQueue) enqueue(el int) {
	b.mu.Lock()
	for b.size == b.capacity {
		<-b.enq
	}

	if b.size == 0 {
		defer func() { b.dq <- 1 }()
	}
	b.tail += 1
	b.data[b.tail] = el
	atomic.AddUint64(&b.size, 1)
	b.mu.Unlock()
}

func (b *BlockingQueue) dequeue() int {
	b.mu.Lock()
	for b.size == 0 {
		<-b.dq
	}

	if b.size == b.capacity {
		defer func() { b.enq <- 1 }()
	}

	el := b.data[b.head]
	b.head += 1
	b.tail--
	b.mu.Unlock()
	return el
}
