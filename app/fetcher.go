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

	url := "https://www.google.com"
	rec := makeFetch(url)

	fmt.Println("** Fetched **")
	fmt.Println(rec)
}

func makeFetch (url string) *FetchRecord {
	// init new record
	nRec := &FetchRecord{Url: url}

	start := time.Now()

	resp, err := http.Get(url)

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

	return nRec
}

