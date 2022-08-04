package main

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

func ValidLink(link string) bool {
	_, e := url.Parse(link)
	if e != nil {
		return false
	}

	if !strings.Contains(strings.ToLower(link), "zippyshare.com") {
		return false // Invalid Link Domain
	}

	//return strings.HasSuffix(link, "file.html")
	return true
}

func LimitRedirects(resp *http.Response, maxRedirectCount int) (*http.Response, error) {

	for i := 0; i < maxRedirectCount; i++ {
		if resp.StatusCode == http.StatusTemporaryRedirect || resp.StatusCode == http.StatusPermanentRedirect {
			redirectLocation := resp.Header.Get("Location")

			//Redirects to nowhere
			if redirectLocation == stringEmpty {
				return nil, errors.New("no redirect location")
			}

			r, e := http.Get(redirectLocation)
			if LogErrorIfNecessary(stringEmpty, &e) {
				return nil, e
			}
			_ = resp.Body.Close()
			resp = r

		} else {
			if resp.Header.Get("Location") != stringEmpty {
				return nil, errors.New("redirect loop")
			} else {
				//Return properly
				return resp, nil
			}
		}
	}
	return nil, errors.New("too many redirects")
}

func GetLinkContent(link string) (*http.Response, error) {
	l, e1 := http.Get(link)
	if LogErrorIfNecessary(stringEmpty, &e1) {
		return nil, e1
	}

	l, e2 := LimitRedirects(l, 5)
	//Only allow OK results to be returned
	if e1 != nil || e2 != nil || l.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(l.StatusCode))
	}

	return l, nil
}
