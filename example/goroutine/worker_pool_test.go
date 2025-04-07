package goroutine

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"testing"
)

func TestDoWork(t *testing.T) {
	input := 42
	expected := sha256.Sum256([]byte(fmt.Sprintf("payload-%d", input)))

	actual := doWork(input)

	t.Logf("hash : %v", actual.String())

	if actual != expected {
		t.Errorf("Expected %x but got %x", expected, actual)
	}
}

func TestWorker(t *testing.T) {
	jobs := make(chan int, 10)
	results := make(chan hashByte, 10)

	for w := 1; w <= numWorkers; w++ {
		t.Logf("Starting worker %d", w)
		go worker(w, jobs, results)
	}

	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range inputs {
		jobs <- v
	}
	close(jobs)

	expected := make(map[int]hashByte)
	for _, v := range inputs {
		expected[v] = doWork(v)
	}

	received := make(map[hashByte]bool)
	for i := 0; i < len(inputs); i++ {
		hashValue := <-results
		received[hashValue] = true
	}

	for _, exp := range expected {
		if !received[exp] {
			t.Errorf("Expected hash %x not received", exp)
		}
	}
}

func BenchmarkUnboundedGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(numJobs)

		for j := 0; j < numJobs; j++ {
			go func(job int) {
				_ = doWork(job)
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(numJobs)

		jobs := make(chan int, numJobs)
		for w := 0; w < numWorkers; w++ {
			go workerBench(jobs, &wg)
		}

		for j := 0; j < numJobs; j++ {
			jobs <- j
		}

		close(jobs)
		wg.Wait()
	}
}
