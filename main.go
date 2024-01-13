package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Write your go routines here  ..
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
