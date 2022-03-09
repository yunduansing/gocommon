package workqueue

// Submit submit new task to the queue
func Submit(task func()) {
	worker.submit(task)
}
