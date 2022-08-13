package Iterations

import (
	"github.com/Arion-Kun/GoZippy/FragmentVariants/Utilities"
	"net/http"
	"regexp"
	"strconv"
)

type Result3 struct{}

func (z Result3) TryParse(response *http.Response) (bool, *Utilities.ZippyFile) {

	bodyPtr := readBody(response)
	scriptPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonSubStr)
	if scriptPtr == nil {
		return false, nil
	}

	fragment := getVersion3Fragments(bodyPtr)
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

const constantLocalVariableDeclaration = 742589

var getVersion3Fragments = func(bodyPtr *string) *[]string {

	gen := GetLinkGeneratorRegex3()
	scriptContentPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonSubStr)
	linkFragments := gen.FindStringSubmatch(*scriptContentPtr)
	if len(linkFragments) != 5 { // 0 is the whole string, 1 is the id, 2 is the key, 3 is the encoded name
		return nil
	}

	// /d/[id]/[key]/[encoded name]
	//[0] is the whole string, [1] is the variable, [2] is the id, [3] is the additional variable, [4] is the encoded name

	returnArray := make([]string, 3)
	returnArray[0] = linkFragments[2] // id
	var_a, _ := strconv.ParseInt(linkFragments[1], 10, 64)
	additional_var, _ := strconv.ParseInt(linkFragments[3], 10, 64)
	key := var_a + additional_var

	returnArray[1] = strconv.FormatInt(key%constantLocalVariableDeclaration, 10) // key
	returnArray[2] = linkFragments[4]                                            // encodedName

	return &returnArray
}

const LINK_GENERATOR_REGEX3 = "var a = (\\d+);[\\s\\S]+?document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\+\\(.+?(\\d+).+?\"/([/\\w%.-]+)\";?"

var rLinkGenerator3 *regexp.Regexp

func GetLinkGeneratorRegex3() *regexp.Regexp {

	if rLinkGenerator3 == nil {
		rLinkGenerator3 = regexp.MustCompile(LINK_GENERATOR_REGEX3)
	}
	return rLinkGenerator3
}
