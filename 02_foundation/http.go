package main

import (
	"net/http"
	"io"
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now().UnixNano()
	var ln int
	wg.Add(3)
	go responseSize("https://www.google.com/", &wg, &ln)
	go responseSize("https://www.google.com/", &wg, &ln)
	go responseSize("https://www.google.com/", &wg, &ln)
	wg.Wait()
	end := time.Now().UnixNano();
	fmt.Println(ln, (end-start))
}

func responseSize(url string, wg *sync.WaitGroup, ln *int) int {
	defer wg.Done()
	response, _ := http.Get(url)
	defer response.Body.Close();
	body,_ := io.ReadAll(response.Body)
	*ln += len(body);
	return *ln
}
