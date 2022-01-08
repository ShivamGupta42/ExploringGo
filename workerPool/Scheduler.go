package workerPool

import (
	"context"
	"sync"
	"time"
)

// ScheduledJob function
type ScheduledJob func(ctx context.Context)

// Scheduler
type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

// Init
func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

// Scheduler Job Addition
func (s *Scheduler) Add(ctx context.Context,
	job ScheduledJob,
	interval time.Duration,
	active bool,
) (chan bool, chan bool) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)

	triggerChannel := make(chan bool)
	activeChannel := make(chan bool)

	s.wg.Add(1)

	// start new thread
	go s.process(ctx,
		job,
		interval,
		triggerChannel,
		activeChannel,
		active,
	)
	return triggerChannel, activeChannel
}

func (s *Scheduler) Stop() {
	//send close signal
	for _, cancel := range s.cancellations {
		cancel()
	}
	// wait for the scheduler finished their job
	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context,
	job ScheduledJob,
	interval time.Duration,
	trigger chan bool,
	activeCh chan bool,
	active bool,
) {
	ticker := time.NewTicker(interval)
	first := make(chan bool, 1)
	first <- true
	isActive := active

	run := func() {
		if isActive {
			job(ctx)
		}
	}

	for {
		select {
		case a := <-activeCh: //changing open status
			isActive = a
		case <-first: // for immediately starting
			run()
		case <-ticker.C: // Scheduler interval occurs - daily routine
			run()
		case <-trigger: // open night shift
			run()
			<-ticker.C //setting daily routine to done it's optional.                          //       you may not work in the morning if night shift tooks to
			//       long and daily routine is close
		case <-ctx.Done(): // close signal
			// closing factory bankrupt
			s.wg.Done()
			ticker.Stop()
			close(first)
			return
		}
	}
}
