package workqueue

import (
	"context"
	"github.com/yunduansing/gocommon/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var (
	worker *workQueue
)

// Start start worker
func Start(ctx context.Context) {
	summary := "worker"
	logger.Info(summary, zap.String("msg", "start worker..."))
	worker = newWorkQueue()
	go worker.run()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(ctx)
	select {
	case <-ctx.Done():
		logger.Info(summary, zap.String("info", "work queue terminating by context canceled"))
	case <-sig:
		logger.Info(summary, zap.String("info", "work queue terminating via signal"))
	}
	cancel()
	worker.close()
}

// Submit submit new task to the queue
func Submit(task func()) {
	worker.submit(task)
}
