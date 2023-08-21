// manager accepts requests from user to start and stop tasks
// manager is responsible for managing the task queue
// manager keep track of the state of the tasks
package manager

import (
	"github.com/okazaki-kk/ayaborg/queue"
	"github.com/okazaki-kk/ayaborg/task"
	"github.com/okazaki-kk/ayaborg/worker"
)

type Manager struct {
	PendingTasks queue.Queue
	TaskDB       map[string]task.Task
	EventDB      map[string]task.TaskEvent
	Workers      []worker.Worker
}

func (m *Manager) SelectWorker() {}
func (m *Manager) UpdateTasks()  {}
func (m *Manager) SendWork()     {}
