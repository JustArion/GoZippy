package Iterations

import (
	"fmt"
	"github.com/Arion-Kun/GoZippy/FragmentVariants/Utilities"
	"github.com/Knetic/govaluate"
	"net/http"
	"regexp"
	"strconv"
)

type Result0 struct{}

func (z Result0) TryParse(response *http.Response) (bool, *Utilities.ZippyFile) {
	bodyPtr := readBody(response)
	scriptPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonSubStr)
	if scriptPtr == nil {
		return false, nil
	}

	fragment := getLinkFragment(bodyPtr)
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

var getLinkFragment = func(bodyPtr *string) *[]string {

	a := GetLinkGeneratorRegex0()
	match := a.FindStringSubmatch(*bodyPtr)
	if len(match) != 4 {
		return nil
	}
	returnArray := make([]string, 3)
	returnArray[0] = match[1] // id
	expr, e := govaluate.NewEvaluableExpression(match[2])
	if e != nil {
		//eMsg3 := fmt.Sprintf("could not create expression: %s", e.Error())
		return nil
	}

	//Solve rawKey
	exp, e1 := expr.Evaluate(nil) // No parameters are passed to the expression
	if e1 != nil {
		//eMsg3 := fmt.Sprintf("could not create expression: %s", e.Error())
		return nil
	}

	var solvedExpression = fmt.Sprintf("%v", exp)

	returnArray[1] = solvedExpression // key
	returnArray[2] = match[3]         // encodedName

	return &returnArray
}

const SCRIPT_REGEX = "(?s)<script type=\"text/javascript\">(.+?)</script>"
const LINK_GENERATOR_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\s*\\+\\s*([\\d\\w\\s+\\-*/%()]+?)\\s*\\+\\s*\"/([/\\w%.-]+)\";?"

var rScript0 *regexp.Regexp

func GetScriptRegex0() *regexp.Regexp {

	if rScript0 == nil {
		rScript0 = regexp.MustCompile(SCRIPT_REGEX)
	}
	return rScript0
}

var rLinkGenerator0 *regexp.Regexp

func GetLinkGeneratorRegex0() *regexp.Regexp {

	if rLinkGenerator0 == nil {
		rLinkGenerator0 = regexp.MustCompile(LINK_GENERATOR_REGEX)
	}
	return rLinkGenerator0
}
