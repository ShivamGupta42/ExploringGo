package web

import "fmt"

type Job struct {
	endInt int
}

type Pool struct {
	jobQueue  chan Job      // Prodcution Line
	readyPool chan chan Job // Worker Channel
}

// Manager in action
func (q *Pool) dispatch() {
	for {
		select {
		case job := <-q.jobQueue:
			workerXChannel := <-q.readyPool //free worker x founded
			workerXChannel <- job           // here is your job worker x
		}
	}
}

type worker struct {
	id int
	readyPool chan chan Job //get work from the boss
	job       chan Job
}

func process(job Job, w *worker) {
	for i := 0; i < job.endInt; i++ {
		fmt.Printf("%d : This is the %d iteration\n", w.id, i)
	}
}

func (w *worker) Start() {

	for {
		w.readyPool <- w.job //hey i am ready to work on new job
		select {
		case job := <-w.job: // hey i am waiting for new job
			process(job,w) // ok i am on it
		}
	}

}
