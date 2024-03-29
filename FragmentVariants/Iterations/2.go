package Iterations

import (
	"bytes"
	"fmt"
	"github.com/Arion-Kun/GoZippy/FragmentVariants/Utilities"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type Result2 struct{}

func (z Result2) TryParse(response *http.Response) (bool, *Utilities.ZippyFile) {

	bodyPtr := readBody(response)
	scriptPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonSubStr)
	if scriptPtr == nil {
		return false, nil
	}

	fragment := getVersion2Fragments(bodyPtr)
	if fragment == nil {
		return false, nil
	}

	ZippyFile := &Utilities.ZippyFile{}
	//Domain is response.Request.URL.Host
	ZippyFile.Domain = response.Request.URL.Host
	ZippyFile.Id = (*fragment)[0]
	_key, err := strconv.ParseInt((*fragment)[1], 10, 64)
	if err != nil {
		return false, nil
	}
	ZippyFile.Key = _key
	ZippyFile.EncodedName = (*fragment)[2]

	//Get Domain:
	//s := response.Request.URL
	//fmt.Println(s.Host)

	return true, ZippyFile
}

const dlbuttonSubStr = "document.getElementById('dlbutton')"
const dlbuttonOmgSubStr = dlbuttonSubStr + ".omg"

var getVersion2Fragments = func(bodyPtr *string) *[]string {

	gen := GetLinkGeneratorRegex2()
	scriptContentPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonSubStr)
	linkFragments := gen.FindStringSubmatch(*scriptContentPtr)
	if len(linkFragments) != 4 { // 0 is the whole string, 1 is the id, 2 is the key, 3 is the encoded name
		return nil
	}

	returnArray := make([]string, 3)
	returnArray[0] = linkFragments[1] // id
	key, _ := strconv.ParseInt(linkFragments[2], 10, 64)

	returnArray[1] = strconv.FormatInt(key%1000+11, 10) // key
	returnArray[2] = linkFragments[3]                   // encodedName

	return &returnArray
}

var readBody = func(rc *http.Response) *string {

	buf := new(bytes.Buffer)
	// After reading the body, the http.Response.Body field is always nil.
	// Therefore, we need to create a new buffer to store the body.

	body, err := ioutil.ReadAll(rc.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	rc.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	buf.Write(body)
	newStr := buf.String() // The variable pointer is now a string

	return &newStr
}

const LINK_GENERATOR_REGEX2 = `document\.getElementById\('dlbutton'\)\.href\s*=\s*"/d/(\w+)/"\+\(([^\%]+)\%1000\s*\+[^\"]+"/(.+?)";?`

var rLinkGenerator2 *regexp.Regexp

func GetLinkGeneratorRegex2() *regexp.Regexp {

	if rLinkGenerator2 == nil {
		rLinkGenerator2 = regexp.MustCompile(LINK_GENERATOR_REGEX2)
	}
	return rLinkGenerator2
}
