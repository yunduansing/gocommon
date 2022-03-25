package workqueue

import (
	"context"
	"github.com/yunduansing/gocommon/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type workQueue struct {
	task    chan func()
	size    int
	cap     int
	running int
	wg      sync.WaitGroup
	cond    sync.Cond
}

func (w *workQueue) submit(t func()) {
	w.task <- t
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
	return &workQueue{task: make(chan func())}
}

var (
	worker *workQueue
)

// Start 开启定时任务
func Start(ctx context.Context) {
	logger.Info("start work...")
	worker = newWorkQueue()
	go worker.run()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ctx.Done():
		logger.Info("", zap.String("info", "work queue terminating by context canceled"))
	case <-sig:
		logger.Info("", zap.String("info", "work queue terminating via signal"))
	}
	worker.close()
}
