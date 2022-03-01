package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Download struct {
	link              string
	fileName          string
	destinationFolder string
}

func NewDownload(link string, destinationFolder string) (*Download, error) {

	_, destErr := os.Stat(destinationFolder)

	//If the folder doesn't exist, return null
	if os.IsNotExist(destErr) {
		return nil, destErr
	}

	escapedLink, err := url.QueryUnescape(path.Base(link))
	if err != nil {
		return nil, err
	}
	decodedLink := escapedLink

	dl := &Download{}
	dl.link = link
	dl.fileName = decodedLink
	dl.destinationFolder = destinationFolder

	return dl, nil
}

func (download *Download) DownloadFile() bool {
	if !Silent {
		fmt.Printf("%sDownloading file: '%s' to %s'%s'\n", blue, download.fileName, reset, download.destinationFolder)
	}

	req, _ := http.NewRequest("GET", download.link, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		LogErrorIfNecessary("", &err)
		return false
	}

	defer func(Body io.ReadCloser) {
		e1 := Body.Close()
		if e1 != nil {
			LogErrorIfNecessary("", &err)
		}
	}(resp.Body)

	if LogErrorIfNecessary(fmt.Sprintf("Failed to download %s to %s: %s", download.link, download.destinationFolder, err), &err) {
		return false
	}

	partialFileFormat := fmt.Sprintf("%s%s", download.fileName, ".tmp") // it is .tmp but it can be any extension )
	tempFolder := path.Join(os.TempDir(), partialFileFormat)
	f, e2 := os.OpenFile(tempFolder, os.O_CREATE|os.O_WRONLY, 0644)

	// We don't care if this errors, as long as the file is gone if something dramatic happens such as a panic
	// Unfortunately the file stagnates in the Temp Folder is the application is interrupted with things like CTRL+C (^C)
	// Or a general computer power failure.
	// New data won't be appended, it is newly written to, so file stagnation leading to file corruption won't be a problem
	defer os.Remove(tempFolder)

	if LogErrorIfNecessary(fmt.Sprintf("Failed to create file %s: %s", download.fileName, e2), &e2) {
		return false
	}

	var err2 error
	if Silent {
		_, e3 := io.Copy(f, resp.Body)
		err2 = e3
	} else {
		bar := CreateProgressBar(resp.ContentLength)

		_, e3 := io.Copy(io.MultiWriter(f, bar), resp.Body)
		err2 = e3
	}
	e4 := f.Close()
	if e4 != nil {
		_ = os.Remove(tempFolder) // Cleanup trash
		return false
	}
	e5 := os.Rename(tempFolder, path.Join(download.destinationFolder, download.fileName))
	if e5 != nil {
		LogErrorIfNecessary("Unable to rename", &e5)
		return false
	}

	if LogErrorIfNecessary(fmt.Sprintf("Failed to copy file %s: %s", download.fileName, err2), &err2) {
		return false
	}

	if !Silent {
		fmt.Printf("%sComplete: %s%s\n", blue, download.fileName, reset)
	}
	return true
}

func CreateProgressBar(maxBytes int64) *progressbar.ProgressBar {
	desc := "Downloading"

	bar := progressbar.DefaultBytes(maxBytes, desc)
	return bar
}

func TryDownload(link string) {
	if !download {
		return
	}

	downloadPtr, err := NewDownload(link, cachedFolderLocation)
	if err != nil {
		LogErrorIfNecessary(fmt.Sprintf("Skipping Download of %s", link), &err)
		return
	}

	currentLink := downloadPtr.link

	escapedLink, err := url.QueryUnescape(path.Base(currentLink))
	if err != nil {
		LogErrorIfNecessary(fmt.Sprintf("Skipping Download of %s", currentLink), &err)
		return
	}

	//LogProgressIfNecessary(fmt.Sprintf("Starting download of %s", escapedLink))
	if !downloadPtr.DownloadFile() { // Finish is already logged inside the download function
		LogProgressIfNecessary(fmt.Sprintf("Failed to download %s", escapedLink))
	}
}

func LogProgressIfNecessary(progress string) {
	if !Silent {
		fmt.Println(progress)
	}
}
