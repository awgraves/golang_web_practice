package main

import (
	"fmt"
	"net/http"
	"time"
)

type FetchRecord struct {
	Url string
	Successful bool
	StatusCode uint8
	ExecTime uint32  // milliseconds
}

func (r *FetchRecord) String () string {
	return fmt.Sprintf("Url: %v\nStatus code: %v\nExecution time: %d milliseconds\n",
						r.Url,
						r.StatusCode,
						r.ExecTime)
}

func fetchMenu() {
	fmt.Print("Make a GET request?\n\n")

	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.codecademy.com",
	}

	// setup buffered channel
	resultChan := make(chan *FetchRecord, len(urls))
	defer close(resultChan)

	// spawn go routines
	for i := range urls {
		url := urls[i]
		go makeFetch(url, resultChan)
	}

	// express results
	for i := 0; i < len(urls); i++ {
		rec := <-resultChan
		fmt.Println("** Fetched **")
		fmt.Println(rec)
	}
}

func makeFetch (url string, resultChan chan *FetchRecord) {
	// init new record
	nRec := &FetchRecord{Url: url}

	// begin timer
	start := time.Now()
	// make request
	resp, err := http.Get(url)
	// stop timer
	elapsed := time.Since(start)

	// calc and convert elapsed time
	timeElapsed := uint32(elapsed.Milliseconds())
	nRec.ExecTime = timeElapsed

	nRec.StatusCode = uint8(resp.StatusCode)

	if err != nil {
		fmt.Println(err)
	} else {
		nRec.Successful = true
	}

	resultChan <- nRec
}

