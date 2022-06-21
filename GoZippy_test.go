package main

import (
	"fmt"
	"os"
	"testing"
)

var Test_Links = []string{
	"https://www120.zippyshare.com/v/PWRXlkLH/file.html",
	"https://www3.zippyshare.com/v/CDCi2wVT/file.html",
}

func Test_SiteConnection(t *testing.T) {
	var success = false
	for _, link := range Test_Links {
		_, err1 := GetLinkContent(link)
		if err1 == nil {
			success = true
			return
		}
	}
	if success == false {
		t.Error("All links failed the test, the site seems offline.")
	}
}

func Test_MakeZippyFile(t *testing.T) {
	var success = false
	var linkHolder = ""
	for _, link := range Test_Links {
		sitePtr, err1 := GetLinkContent(link)
		if err1 != nil {
			continue
		}

		filePtr, successful := TryMakeZippyFile(sitePtr)
		if successful {
			success = true
			cachedFile = filePtr
			return
		} else {
			linkHolder = link
		}
	}
	if success == false {
		t.Error(fmt.Sprintf("Failed to make zippy file for %s", linkHolder))
	}
}

var cachedFile *ZippyFile

func Test_DownloadFile(t *testing.T) {

	if cachedFile == nil {
		t.Error("MakeZippyFile failed, cannot test download.")
		return
	}

	cachedFolderLocation = "."
	download = true
	downloadPath := TryDownload(cachedFile.GetEncodedLink())
	if downloadPath == nil {
		return
	}
	//  Remove TryDownload Artifacts
	e := os.Remove(*downloadPath)
	if e != nil {
		t.Error("Failed to remove download artifact", *downloadPath)
	} else {
		t.Log("Removed download artifact from", *downloadPath)
	}
}
