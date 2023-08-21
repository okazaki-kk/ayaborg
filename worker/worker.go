// worker run tasks as docker containers
// accept requests from manager to start and stop tasks
package worker

import (
	"github.com/okazaki-kk/ayaborg/queue"
	"github.com/okazaki-kk/ayaborg/task"
)

type Worker struct {
	Name      string
	Db        map[string]task.Task
	Queue     queue.Queue
	TaskCount int
}

func (w *Worker) RunTask()  {}
func (w *Worker) StopTask() {}
