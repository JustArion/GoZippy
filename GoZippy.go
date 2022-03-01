//Inspired from: https://github.com/scotow/zippyst

package main

import (
	"bufio"
	"fmt"
	"github.com/Arion-Kun/GoLaunch"
	"log"
	"os"
	"sync"
)

func main() {

	// 1 is the program just being run without launch args since 1 is the program path + name
	if len(os.Args) == 1 {
		printHelp()
		return
	}
	initializeStartup()
}

const stringEmpty = ""

var links []string

func getFile(launchArg string) {
	markDirty = true

	filePointer, err := os.Open(GoLaunch.Get(launchArg)[0])
	if err != nil && !Silent {
		log.Fatal("Unable to open file: ", err)
	}
	scanner := bufio.NewScanner(filePointer)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}
	if er := scanner.Err(); er != nil && !Silent {
		log.Fatal("Error reading file: ", er)
	}

	// If the file fails to close log the error.
	defer func(filePointer *os.File) {
		e := filePointer.Close()
		if e != nil && !Silent {
			log.Fatal("Error closing file: ", e)
		}
	}(filePointer)
}

func getLinks(launchArg string) {
	markDirty = true

	containsValue, value := GoLaunch.TryGetValue(launchArg)

	if !containsValue {
		if !Silent {
			log.Fatalln("ERROR: No value found for the launch argument: " + launchArg)
		}
		return
	}

	if value == nil {
		if !Silent {
			println("No link specified. Please follow -L or --link with a link. (Example: ./GoZippy.exe -L 'https://www3.zippyshare.com/v/CDCi2wVT/file.html' )")
		}
	}

	links = append(links, value...)

}

// InspectLink Inspect Process:
// Validate Link
// Parse Link
// Print Link to Output / stdout
// Download Link
func InspectLink(link string, awaiter *sync.WaitGroup) {
	if awaiter != nil {
		defer awaiter.Done()
	}
	if !ValidLink(link) && !Silent {
		log.Fatalln("Invalid link: " + link)
		return
	}
	rawWebsitePtr, e1 := GetLinkContent(link)
	if LogErrorIfNecessary("Error getting link content:", &e1) {
		return
	}
	file, made := TryMakeZippyFile(rawWebsitePtr)
	if !made {
		return
	}
	outputLink := file.GetEncodedLink()
	fmt.Println(outputLink)
	if awaiter == nil {
		TryDownload(outputLink)
	}

}

// DownloadBuffer I would rather split the buffers than have a single buffer with input links being overridden with output links.
var DownloadBuffer []string

func InspectLinkAndSort(link string, index int, awaiter *sync.WaitGroup) {
	defer awaiter.Done()

	if !ValidLink(link) && !Silent {
		log.Fatalln("Invalid link" + link)
		return
	}
	rawWebsitePtr, e1 := GetLinkContent(link)
	if LogErrorIfNecessary("Error getting link content", &e1) {
		return
	}
	file, made := TryMakeZippyFile(rawWebsitePtr)
	if !made {
		return
	}
	DownloadBuffer[index] = file.GetEncodedLink()

}

func LogErrorIfNecessary(errorMessage string, err *error) bool {
	if *err != nil && !Silent {
		fmt.Printf("%s: %s\n", errorMessage, *err)
		return true
	}
	return false
}
