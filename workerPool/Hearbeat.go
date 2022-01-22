package workerPool

type Pool struct {
	jobQueue  chan Job      // Prodcution Line
	readyPool chan chan Job // Worker Channel
}

func (q *Pool) Submit(j *Job) {

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
	readyPool chan chan Job //get work from the boss
	job       chan Job
}

func (w *worker) Start() {

	for {
		w.readyPool <- w.job //hey i am ready to work on new job
		select {
		case job := <-w.job: // hey i am waiting for new job
			w.Process(job, w) // ok i am on it
		}
	}

}

func (w *worker) Process(job Job, w2 *worker) {

}
