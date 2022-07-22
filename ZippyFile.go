package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ZippyFile struct {
	domain      string
	id          string
	key         int64
	encodedName string
}

func (wf *ZippyFile) GetLink() string {
	return fmt.Sprintf("https://%s/d/%s/%d/DOWNLOAD", wf.domain, wf.id, wf.key)
}
func (wf *ZippyFile) GetEncodedLink() string {
	return fmt.Sprintf("https://%s/d/%s/%d/%s", wf.domain, wf.id, wf.key, wf.encodedName)
}

func TryMakeZippyFile(response *http.Response) (*ZippyFile, bool) {

	bodyPtr := readBody(response)
	scriptPtr := getScriptContent(bodyPtr)
	if scriptPtr == nil {
		return nil, false
	}

	//TEST
	//vars := GetVariableRegex().FindAllStringSubmatch(*scriptPtr, -1)
	//if vars == nil && !Silent {
	//	fmt.Println("var: ", vars)
	//	return nil, false
	//}
	//

	fragment := getLinkFragments(bodyPtr)
	if fragment == nil {
		return nil, false
	}

	ZippyFile := &ZippyFile{}
	//Domain is response.Request.URL.Host
	ZippyFile.domain = response.Request.URL.Host
	ZippyFile.id = (*fragment)[0]
	_key, err := strconv.ParseInt((*fragment)[1], 10, 64)
	if err != nil {
		return nil, false
	}
	ZippyFile.key = _key
	ZippyFile.encodedName = (*fragment)[2]

	//Get Domain:
	//s := response.Request.URL
	//fmt.Println(s.Host)

	return ZippyFile, true
}

const dlbuttonSubStr = "document.getElementById('dlbutton')"

var getScriptContent = func(bodyPtr *string) *string {

	matches := GetScriptRegex().FindAllString(*bodyPtr, 15)
	for i := range matches {
		if strings.Contains(matches[i], dlbuttonSubStr) {
			return &matches[i]
		}
	}
	eMsg1 := "dlButtonSubStr could not be found"
	e := errors.New(eMsg1)
	LogErrorIfNecessary(eMsg1, &e)
	return nil
}

var getLinkFragments = func(bodyPtr *string) *[]string {

	gen := GetLinkGeneratorRegex()
	scriptContentPtr := getScriptContent(bodyPtr)
	linkFragments := gen.FindStringSubmatch(*scriptContentPtr)
	if len(linkFragments) != 4 { // 0 is the whole string, 1 is the id, 2 is the key, 3 is the encoded name
		eMsg2 := fmt.Sprintf("link fragment length does not match: %d", len(linkFragments))
		e := errors.New(eMsg2)
		LogErrorIfNecessary(eMsg2, &e)
		return nil
	}

	returnArray := make([]string, 3)
	returnArray[0] = linkFragments[1] // id
	key, _ := strconv.ParseInt(linkFragments[2], 10, 64)

	returnArray[1] = strconv.FormatInt(key%1000 + 11, 10) // key
	returnArray[2] = linkFragments[3]                 // encodedName

	return &returnArray
}
