package main

import (
	"errors"
	"fmt"
	"math"
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
	if vars == nil && !Silent {
		fmt.Println("var: ", vars)
		return nil, false
	}
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
	if len(linkFragments) != 3 { // 0 is the whole string, 1 is the id, 2 is the encoded name
		eMsg2 := fmt.Sprintf("link fragment length does not match: %d", len(linkFragments))
		e := errors.New(eMsg2)
		LogErrorIfNecessary(eMsg2, &e)
		return nil
	}

	//match := gen.FindStringSubmatch(*bodyPtr)
	//if len(match) != 2 {
	//	eMsg2 := fmt.Sprintf("link fragment length does not match: %d", len(match))
	//	e := errors.New(eMsg2)
	//	LogErrorIfNecessary(eMsg2, &e)
	//	return nil
	//}
	returnArray := make([]string, 3)
	returnArray[0] = linkFragments[1] // id

	// Not needed yet but for the future, since it seems like the owner wants to automate this crap.
	//variableKey := GetVariableKeyRegex().FindStringSubmatch(*bodyPtr)           // "asdasd"
	//if variableKey == nil || len(variableKey) != 2 || len(variableKey[1]) < 3 { // We look for the variable key and  substring it so there needs to be at least 1 match and its length must be minimum 3
	//	eMsg3 := "variableKey could not be found"
	//	e := errors.New(eMsg3)
	//	LogErrorIfNecessary(eMsg3, &e)
	//	return nil
	//}

	variable := GetVariableRegex().FindStringSubmatch(*scriptContentPtr)

	// Substring the first capture from the variable key by 3

	a, e := strconv.ParseFloat(variable[1], 64)
	if e != nil {
		eMsg5 := "variable could not be parsed to int"
		e := errors.New(eMsg5)
		LogErrorIfNecessary(eMsg5, &e)
		return nil
	}

	// Substring from 0 -> 3 from the string variable
	key := math.Pow(a, 2) + 30
	returnArray[1] = strconv.FormatFloat(key, 'f', 0, 64)

	//expr, e := govaluate.NewEvaluableExpression(match[2])
	//if e != nil {
	//	eMsg3 := fmt.Sprintf("could not create expression: %s", e.Error())
	//	e := errors.New(eMsg3)
	//	LogErrorIfNecessary(eMsg3, &e)
	//	return nil
	//}

	//Solve rawKey
	//exp, e1 := expr.Evaluate(nil) // No parameters are passed to the expression
	//if e1 != nil {
	//	eMsg3 := fmt.Sprintf("could not create expression: %s", e.Error())
	//	LogErrorIfNecessary(eMsg3, &e1)
	//	return nil
	//}
	returnArray[2] = linkFragments[2] // encodedName

	return &returnArray
}
