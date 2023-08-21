// task is the smallest unit of the orchestration system
// it is a wrapper around a container
// it is responsible for the container lifecycle and state and restart-policy
package task

type State int

const (
	Running State = iota
	Stopped
	Pending
	Scheduled
	Failed
)

type Task struct {
	ID            string
	Name          string
	State         State
	Image         string
	RestartPolicy string
	Memory        int
	Disk          int
}

// TaskEvent represents an event that moves a task from one state to another
type TaskEvent struct {
	ID    string
	State State
	Task  Task
}
