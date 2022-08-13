package Iterations

import (
	"fmt"
	"github.com/Arion-Kun/GoZippy/FragmentVariants/Utilities"
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

	unsolvedExpression := match[2]
	expressionValues := GetEvaluationRegex0().FindStringSubmatch(unsolvedExpression)
	if len(expressionValues) != 5 {
		return nil
	}
	// (426765 % 51245 + 426765 % 913)
	val1, _ := strconv.ParseInt(expressionValues[1], 10, 64)
	val2, _ := strconv.ParseInt(expressionValues[2], 10, 64)
	val3, _ := strconv.ParseInt(expressionValues[3], 10, 64)
	val4, _ := strconv.ParseInt(expressionValues[4], 10, 64)

	exp := val1%val2 + val3%val4

	var solvedExpression = fmt.Sprintf("%v", exp)

	returnArray[1] = solvedExpression // key
	returnArray[2] = match[3]         // encodedName

	return &returnArray
}

const SCRIPT_REGEX = "(?s)<script type=\"text/javascript\">(.+?)</script>"
const EVALUATION_REGEX = "\\((\\d+) % (\\d+) \\+ (\\d+) % (\\d+)\\)"
const LINK_GENERATOR_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\s*\\+\\s*([\\d\\w\\s+\\-*/%()]+?)\\s*\\+\\s*\"/([/\\w%.-]+)\";?"

var rEvaluation0 *regexp.Regexp

func GetEvaluationRegex0() *regexp.Regexp {

	if rEvaluation0 == nil {
		rEvaluation0 = regexp.MustCompile(EVALUATION_REGEX)
	}
	return rEvaluation0
}

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
