package main

import (
	"fmt"
	"testing"
)

var Test_Links = []string{
	"https://www3.zippyshare.com/v/CDCi2wVT/file.html",
	"https://www101.zippyshare.com/v/nMUOpguX/file.html",
}

func TestSiteConnection(t *testing.T) {
	for _, link := range Test_Links {
		_, err1 := GetLinkContent(link)
		if err1 != nil {
			t.Error(fmt.Sprintf("Site is offline on link: %s", link))
		}
	}
}

func TestMakeFileZippyFile(t *testing.T) {

	for _, link := range Test_Links {
		sitePtr, err1 := GetLinkContent(link)
		if err1 != nil {
			t.Error(fmt.Sprintf("Site is offline on link: %s", link))
		}

		_, successful := TryMakeZippyFile(sitePtr)
		if !successful {
			t.Error(fmt.Sprintf("Failed to make zippy file for %s", link))
		}
	}

}
