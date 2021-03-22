package main

import (
	"context"
	"fmt"
	"strings"
)

const WORKER_POOL_SIZE = 3

func worker(cancel context.CancelFunc, ctx context.Context, id int, jobs <-chan string, results chan<- string, errors chan<- error) {
	for {

		select {
		case <-ctx.Done():
			fmt.Println("Done now")
			return
		default:
		}

		job, ok := <-jobs

		if !ok {
			return
		}

		fmt.Printf("Job: %s processed by worker: %d\n", job, id)

		if job != "fort" {
			results <- strings.ToUpper(job)
		} else {
			errors <- fmt.Errorf("Not a valid city: %s", job)
			cancel()
		}
	}
}

func main() {
	arr := []string{"deakin", "lancaster", "kent", "ulster", "fort", "devon", "wrexham", "earley", "hull", "rhyl", "corby"}

	size := len(arr)
	jobs := make(chan string, size)

	results := make(chan string)
	errors := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < WORKER_POOL_SIZE; i++ {
		go worker(cancel, ctx, i, jobs, results, errors)
	}

	for _, item := range arr {
		jobs <- item
	}

	var final []string
	var errorList []string

	received := 0
	for {
		select {
		case <-ctx.Done():
			received = size
			break
		case res := <-results:
			final = append(final, res)
			received++
		case err := <-errors:
			errorList = append(errorList, err.Error())
			received++
		default:
		}
		if received == size {
			break
		}
	}

	fmt.Printf("Results: %v\n", final)
	fmt.Printf("Errors: %v\n", errorList)
}
