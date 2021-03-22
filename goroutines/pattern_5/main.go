package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/10",
		"https://jsonplaceholder.typicode.com/posts/101010",
	}
	for _, url := range urls {
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
				if resp.StatusCode != 200 {
					return fmt.Errorf("%s failed", url)
				}
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	err := g.Wait()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully fetched all URLs.")
	}
}
