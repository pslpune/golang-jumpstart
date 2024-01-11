package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Parallel - what do you mean by that ? Core 1 (CPU1) - task A || Core 2 (CPU2) - task B
// Core 2 Duo Core 1 (4 threads) (clocks, Clockd speed 2 Ghz - in a single sec it is capable of taking 2x10^12 instructions , FSB) = Thread 1 -task A || Thread 2 - TaskB || Thread 3 || Thread4
// Being AGNOSTIC of whether it is scheduled on single core or multicore when a language lets you program uniformly its called *asynchrnoy*
func main() {
	fmt.Println("now learning asynchrony in Go")
	// comic, err := DownloadComic(100)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(comic["safe_title"])
	jobs := []int{}
	for i := 100; i < 120; i++ {
		jobs = append(jobs, i)
	}

	// defer close(results)
	// Non blocking operation
	// channel rules

	// Reading empty channels & Writing to full channels is blocking
	// Writing on closed channels results in a panic, exception
	// Reading from nil channels will result in a panic
	_, done := func() (chan map[string]interface{}, chan bool) { // Go routine
		results := make(chan map[string]interface{}, 20)
		done := make(chan bool, 1)
		// var results chan map[string]interface{}
		go func() {
			// defer close(done)
			defer close(results)
			for _, j := range jobs {
				comic, err := DownloadComic(j)
				if err != nil {
					fmt.Printf("error downloading the comic %s", err)
				} else {
					// instead of printing the safe title here lets go ahead to print the same in the main thread
					fmt.Println(comic["safe_title"])
					// results <- comic
				}
			}
			done <- false
		}()
		return results, done
		// while you can read from closed channels, you cannot write into closed channels
		// reading on empty channels is a blocking operation
		// go func() {
		// }()
	}() // IIFE -  immediately invoked function expression
	<-done
	// for comic := range results {
	// 	fmt.Println(comic["safe_title"])
	// }
	fmt.Println("End of the program")
}

func DownloadComic(comicIndex int) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicIndex)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request %s", err)
	}
	cl := &http.Client{
		Timeout: 4 * time.Second,
	}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request over http %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get favorable response from server %d", resp.StatusCode)
	}
	// this is wehre we read the json response payload
	defer resp.Body.Close()
	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid json payload from server %s", err)
	}
	// defer // this is the first time i have used it ! .. will come to this a bit later
	result := map[string]interface{}{}
	err = json.Unmarshal(byt, &result)
	if err != nil {
		return nil, fmt.Errorf("fauled to read payload from server %s", err)
	}
	return result, nil
}
