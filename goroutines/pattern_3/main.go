package main

import (
	"fmt"
	"strings"
)

const WORKER_POOL_SIZE = 3

func worker(id int, jobs <-chan string, results chan<- string, errors chan<- error) {
	for job := range jobs {
		fmt.Printf("Job: %s processed by worker: %d\n", job, id)
		if job != "fort" {
			results <- strings.ToUpper(job)
		} else {
			errors <- fmt.Errorf("Not a valid city: %s", job)
		}
	}
}

func main() {
	arr := []string{"deakin", "lancaster", "kent", "ulster", "fort", "devon", "wrexham", "earley", "hull", "rhyl", "corby"}

	size := len(arr)
	jobs := make(chan string, size)

	results := make(chan string)
	errors := make(chan error)

	for i := 0; i < WORKER_POOL_SIZE; i++ {
		go worker(i, jobs, results, errors)
	}

	for _, item := range arr {
		jobs <- item
	}

	var final []string
	var errorList []string

	received := 0
	for {
		select {
		case res := <-results:
			final = append(final, res)
			received++
		case err := <-errors:
			errorList = append(errorList, err.Error())
			received++
		}
		if len(final)+len(errorList) >= size {
			break
		}
	}

	fmt.Printf("Results: %v\n", final)
	fmt.Printf("Errors: %v\n", errorList)
}
