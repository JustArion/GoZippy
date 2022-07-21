package main

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
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
	vars := GetVariableRegex().FindAllStringSubmatch(*scriptPtr, -1)
	if vars != nil && !Silent {
		fmt.Println("vars: ", vars)
	}
	//

	fragment := getLinkFragment(bodyPtr)
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

var getLinkFragment = func(bodyPtr *string) *[]string {

	a := GetLinkGeneratorRegex()
	match := a.FindStringSubmatch(*bodyPtr)
	if len(match) != 4 {
		eMsg2 := fmt.Sprintf("link fragment length does not match: %d", len(match))
		e := errors.New(eMsg2)
		LogErrorIfNecessary(eMsg2, &e)
		return nil
	}
	returnArray := make([]string, 3)
	returnArray[0] = match[1] // id
	expr, e := govaluate.NewEvaluableExpression(match[2])
	if e != nil {
		eMsg3 := fmt.Sprintf("could not create expression: %s", e.Error())
		e := errors.New(eMsg3)
		LogErrorIfNecessary(eMsg3, &e)
		return nil
	}

	//Solve rawKey
	exp, e1 := expr.Evaluate(nil) // No parameters are passed to the expression
	if e1 != nil {
		eMsg3 := fmt.Sprintf("could not create expression: %s", e.Error())
		LogErrorIfNecessary(eMsg3, &e1)
		return nil
	}

	var solvedExpression = fmt.Sprintf("%v", exp)

	returnArray[1] = solvedExpression // key
	returnArray[2] = match[3]         // encodedName

	return &returnArray
}
