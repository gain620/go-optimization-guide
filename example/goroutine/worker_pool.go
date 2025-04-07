package goroutine

import (
	"crypto/sha256"
	"fmt"
	"runtime"
	"sync"
)

const (
	numJobs = 10000
)

var numWorkers = runtime.NumCPU()

type hashByte [32]byte

func (h hashByte) String() string {
	return fmt.Sprintf("%x", h[:])
}

func workerBench(jobs <-chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		_ = doWork(job)
		wg.Done()
	}
}

func worker(id int, jobs <-chan int, results chan<- hashByte) {
	for j := range jobs {
		results <- doWork(j)
	}
}

func doWork(n int) hashByte {
	data := []byte(fmt.Sprintf("payload-%d", n))
	return sha256.Sum256(data) //
}
