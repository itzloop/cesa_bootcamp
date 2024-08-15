package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// Job represents a work that should be processed by a worker
type Job struct {
	ID       int
	MetaData string
}

func main() {
	// define some command line arguments
	workerCount := flag.Int("wc", 16, "number of workers")
	jobsCount := flag.Int("jc", 100, "number of jobs")
	jobsChanSize := flag.Int("ch", 16, "jobs channel buffer size")
	useWorkerPool := flag.Bool("w", false, "use worker pool instead of processing jobs sequentially")

	// parse the command line arguments passed to the program
	flag.Parse()

	if *useWorkerPool {
		// create the jobs channel with the given channel size
		jobs := make(chan *Job, *jobsChanSize)

		// create a waitgroup to wait for all go-routines to finish
		wg := sync.WaitGroup{}
		for i := 0; i < *workerCount; i++ {
			// for each go-routine, add to the wait group and decrement
			// from it when the worker is finished
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()

				// run each worker
				worker(workerID, jobs)
			}(i)
		}

		// generate some artificial work for the workers
		for i := 0; i < *jobsCount; i++ {
			jobs <- &Job{
				ID:       i,
				MetaData: fmt.Sprintf("this is job number#%d", i+1),
			}
		}

		// clone the channel
		close(jobs)

		// wait for all workers to finish their work
		wg.Wait()
	} else {
		// generate some artificial work sequentially
		for i := 0; i < *jobsCount; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("job\tid: %d\t metadata: %s is done\n", i, fmt.Sprintf("this is job number#%d", i+1))
		}
	}

}

func worker(workerID int, jobsChan <-chan *Job) {
	fmt.Printf("worker %d started\n", workerID)

	// read from jobs channel and do the job :)
	for job := range jobsChan {
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d is processing job\tid: %d\t metadata: %s\n", workerID, job.ID, job.MetaData)
	}

	fmt.Printf("worker %d done\n", workerID)
}
