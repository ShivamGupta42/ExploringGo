package workerPool

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Job struct {
	Id int
	Wg *sync.WaitGroup
}

func (j *Job) Do() {
	time.Sleep(1 * time.Second)
	log.Println(fmt.Sprintf("Job Finished:%d", j.Id))
	j.Wg.Done()
}
func DailyRoutine(ctx context.Context) {
	log.Println("ScheduledJob Started")
	wg := &sync.WaitGroup{}
	jobs := []int{1, 2, 3, 4}
	for _, j := range jobs {
		wg.Add(1)
		fmt.Println(j)
		// see the library
		//Pool.Submit(&Job{Id: j, Wg: &wg})
	}

	//waiting all jobs to be finished
	wg.Wait()

	log.Println("DailyRoutine Finished")
}

/*
scheduler := NewScheduler()
ctx := context.Background()
scheduler.Add(ctx, DailyRoutine, time.Hour*24, true) //active
*/
