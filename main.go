package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	MAX_THREADS = 10
)

func main() {
	jobs := []int{}
	for i := 100; i <= 120; i++ {
		jobs = append(jobs, i)
	}
	jobsStack := make(chan int, len(jobs))
	for _, j := range jobs {
		jobsStack <- j
	}
	close(jobsStack)
	// done := make(chan bool, 1)
	var wg sync.WaitGroup
	results := make(chan map[string]interface{}, 20)

	for i := 0; i < MAX_THREADS; i++ {
		wg.Add(1)
		go func(jobs chan int) {
			defer wg.Done()
			for j := range jobs {
				result, err := DownloadComic(j)
				if err != nil {
					fmt.Println(err)
					return
				}
				results <- result
			}
		}(jobsStack)
	}
	go func() {
		for r := range results {
			fmt.Println(r["safe_title"])
		}
	}()
	/*
		This is when we just make a async job for all the download jobs
		and the results are passed back to the main thread on the channel instead of the thread printing it
	*/
	// go func(done chan bool) {
	// 	for _, j := range jobs {
	// 		result, err := DownloadComic(j)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		results <- result
	// 	}
	// 	done <- true
	// }(done)
	// for r := range results {
	// 	fmt.Println(r)
	// }
	// <-done

	/*
		No
	*/
	// var wg sync.WaitGroup
	// for _, j := range jobs {
	// 	wg.Add(1)
	// 	go func(j int) {
	// 		result, err := DownloadComic(j)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		fmt.Println(result)
	// 		wg.Done()
	// 	}(j)
	// }
	wg.Wait()
	close(results)
	fmt.Println("We are now closing the xkcd task")
}

func DownloadComic(comicIndex int) (map[string]interface{}, error) {

	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicIndex)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request %s", err)
	}
	cl := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing the request %s", err)
	}
	if resp.StatusCode == http.StatusOK {
		// process the payload
		byt, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response payload from xkcd.com %s", err)
		}
		payload := map[string]interface{}{}
		err = json.Unmarshal(byt, &payload)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling payload in response %s", err)
		}
		// once we have payload we then just pass it back to the calling function
		return payload, nil
	}
	// else in this case we have an error
	return nil, fmt.Errorf("unfavorable http code %d", resp.StatusCode)
}
