package Iterations

import (
	"github.com/Arion-Kun/GoZippy/FragmentVariants/Utilities"
	"net/http"
	"regexp"
	"strconv"
)

type Result4 struct{}

func (z Result4) TryParse(response *http.Response) (bool, *Utilities.ZippyFile) {

	bodyPtr := readBody(response)
	scriptPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonSubStr)
	if scriptPtr == nil {
		return false, nil
	}

	fragment := getVersion4Fragments(bodyPtr)
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

const constantLocalVariableDeclaration4 = 35478

var getVersion4Fragments = func(bodyPtr *string) *[]string {

	gen := GetLinkGeneratorRegex4()
	scriptContentPtr := Utilities.TryFindScriptMatch(GetScriptRegex0(), bodyPtr, dlbuttonOmgSubStr)
	linkFragments := gen.FindStringSubmatch(*scriptContentPtr)
	if len(linkFragments) != 3 { // 0 is the whole string, 1 is the id, 2 is the key, 3 is the encoded name
		return nil
	}

	// /d/[id]/[key]/[encoded name]
	//[0] is the whole string regex, [1] is the variable, [2] is the id, [3] is the additional variable, [4] is the encoded name

	returnArray := make([]string, 3)

	returnArray[0] = linkFragments[1]                                         // id
	returnArray[1] = strconv.FormatInt(constantLocalVariableDeclaration4, 10) // key
	returnArray[2] = linkFragments[2]                                         // encodedName

	return &returnArray
}

const LINK_GENERATOR_REGEX4 = `var a = \d+;[\s\S]+?document\.getElementById\('dlbutton'\)\.href\s*=\s*"/d/(\w+)/"\+\(.+?\d+.+?"/([/\w%.-]+)";?`

var rLinkGenerator4 *regexp.Regexp

func GetLinkGeneratorRegex4() *regexp.Regexp {

	if rLinkGenerator4 == nil {
		rLinkGenerator4 = regexp.MustCompile(LINK_GENERATOR_REGEX4)
	}
	return rLinkGenerator4
}
