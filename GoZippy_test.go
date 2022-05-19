package main

import (
	"fmt"
	"testing"
)

var Test_Links = []string{
	"https://www84.zippyshare.com/v/LGzQUf1d/file.html",
	"https://www10.zippyshare.com/v/LnhJvx0M/file.html",
}

func TestSiteConnection(t *testing.T) {
	var success = false
	for _, link := range Test_Links {
		_, err1 := GetLinkContent(link)
		if err1 == nil {
			success = true
		}
	}
	if success == false {
		t.Error("All links failed the test, the site seems offline.")
	}
}

func TestMakeFileZippyFile(t *testing.T) {

	var success = false
	var linkHolder = ""
	for _, link := range Test_Links {
		sitePtr, err1 := GetLinkContent(link)
		if err1 != nil {
			continue
		}

		_, successful := TryMakeZippyFile(sitePtr)
		if successful {
			success = true
		} else {
			linkHolder = link
		}
	}
	if success == false {
		t.Error(fmt.Sprintf("Failed to make zippy file for %s", linkHolder))
	}

}
