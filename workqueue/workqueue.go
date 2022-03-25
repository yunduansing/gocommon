package workqueue

import (
	"math"
	"sync"
	"sync/atomic"
)

var (
	defaultWorkerSize int32 = math.MaxInt32
)

type workQueue struct {
	task    chan func()
	size    int32
	cap     int32
	running int32
	wg      sync.WaitGroup
	cond    sync.Cond
}

type Option struct {
}

func (w *workQueue) submit(t func()) {
	w.task <- t
	atomic.AddInt32(&w.running, 1)
}

func (w *workQueue) run() {
	for t := range w.task {
		go t()
	}
}

func (w *workQueue) close() {
	close(w.task)
}

func newWorkQueue() *workQueue {
	return &workQueue{task: make(chan func()), cap: defaultWorkerSize}
}

func newWorkQueueWith(cap int32, options ...Option) *workQueue {
	return &workQueue{task: make(chan func()), cap: cap}
}
