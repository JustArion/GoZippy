package main

import (
	"fmt"
	"sync"
)

func StartWorker() {

	DownloadBuffer = make([]string, len(links))
	if Async {
		wg := &sync.WaitGroup{}
		if Sort {
			iterateAsynchronouslyAndSort(wg)
		} else {
			iterateAsynchronously(wg)
		}
	} else {
		iterateSynchronously()
	}
}
func iterateSynchronously() {
	for _, link := range links {
		InspectLink(link, nil)
	}
}
func iterateAsynchronously(awaiter *sync.WaitGroup) {
	awaiter.Add(len(links))
	for _, link := range links {
		go InspectLink(link, awaiter)
	}
	awaiter.Wait()
}
func iterateAsynchronouslyAndSort(awaiter *sync.WaitGroup) {

	awaiter.Add(len(links))
	for i, link := range links {
		go InspectLinkAndSort(link, i, awaiter)
	}
	awaiter.Wait()

	// This is a bit bad because it only prints when they're all done instead of as they come in.
	// The down-side to printing as they come in is Race-Conditions
	for i := range DownloadBuffer {
		fmt.Println(DownloadBuffer[i])
	}
}
