package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)


const fetchListFileName = "fetch_list.txt"

// go does not support slice constants!  use func to get around this
func defaultUrls () []string {
	return []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.codecademy.com",
	}
}
type FetchRecord struct {
	Url string
	Successful bool
	StatusCode uint8
	ExecTime uint32  // milliseconds
}

func (r *FetchRecord) String () string {
	return fmt.Sprintf("Url: %v\nSuccess: %v\nStatus code: %v\nExecution time: %d milliseconds\n",
						r.Url,
						r.Successful,
						r.StatusCode,
						r.ExecTime)
}

func fetchMenu() {
	fmt.Print("Make a GET request?\n\n")

	// grab list from the file (if doesnt exist, default list is provided)
	urls := readUrlListFromFile()

	// setup buffered channel
	resultChan := make(chan *FetchRecord, len(urls))
	defer close(resultChan)

	// spawn go routines
	for i := range urls {
		url := urls[i]
		go makeFetch(url, resultChan)
	}

	// express results & save to array
	var results []*FetchRecord

	for i := 0; i < len(urls); i++ {
		rec := <-resultChan
		fmt.Println("** Result **")
		fmt.Println(rec)
		results = append(results, rec)
	}

	exportResultToFile(results)
}

func readUrlListFromFile() []string {
	// init default list
	urls := defaultUrls()
	// attempt to read bytes of file
	fileBytes, err := ioutil.ReadFile(fetchListFileName)

 	if err != nil {
		 // switch to default list (file might not exist)
		fmt.Printf("!!! Could not find file '%v' in this dir\n", fetchListFileName)
		fmt.Print("** Using default url list instead... **\n\n")
		return urls
 	}

	 // trim off the last newline at EOF then split elements by line into slice
 	if fileUrls := strings.Split(strings.Trim(string(fileBytes), "\n"), "\n"); len(fileUrls) != 0 {
		 // make sure file isn't empty
		 urls = fileUrls
	 }
	 
	return urls
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

	if err != nil {
		nRec.Successful = false
	} else {
		nRec.Successful = true
		nRec.StatusCode = uint8(resp.StatusCode)
	}

	resultChan <- nRec
}

func exportResultToFile (recs []*FetchRecord) {
	fmt.Println("Saving results to fetch_results.txt...")

	// create the file
	f, _ := os.Create("fetch_results.txt")
	defer f.Close()

	// give header to file
	loc, _ := time.LoadLocation("America/New_York")
	dt := time.Now().In(loc).Format("Mon Jan 2, 2006 3:04:05 PM")
	
	h := fmt.Sprintf("** Output of program %v **\n", dt)
	stars := fmt.Sprintln(strings.Repeat("*", len(h) - 1))  // last char is new line char, so -1
	fmt.Fprint(f, stars, h, stars)

	for _, v := range recs {
		fmt.Fprint(f, "\n", v)
	}
}
