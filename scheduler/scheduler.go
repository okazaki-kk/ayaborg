// scheduler determines a set of workers to run a task
// scheduler score the candidate workers from best to worst
// scheduler picks the best worker to run the task
package scheduler

type Scheduler interface {
	SelectWorker()
	Score()
	Pick()
}
