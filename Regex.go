package main

import "regexp"

const SCRIPT_REGEX = "(?s)<script type=\"text/javascript\">(.+?)</script>"
const VARIABLE_REGEX = "var\\s*(\\w+)\\s*=\\s*([\\d+\\-*/%]+);?"
const LINK_GENERATOR_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\s*\\+\\s*([\\d\\w\\s+\\-*/%()]+?)\\s*\\+\\s*\"/([/\\w%.-]+)\";?"

var r_Script *regexp.Regexp

func GetScriptRegex() *regexp.Regexp {

	if r_Script == nil {
		r_Script = regexp.MustCompile(SCRIPT_REGEX)
	}
	return r_Script
}

var r_Variable *regexp.Regexp

func GetVariableRegex() *regexp.Regexp {

	if r_Variable == nil {
		r_Variable = regexp.MustCompile(VARIABLE_REGEX)
	}
	return r_Variable
}

var r_LinkGenerator *regexp.Regexp

func GetLinkGeneratorRegex() *regexp.Regexp {

	if r_LinkGenerator == nil {
		r_LinkGenerator = regexp.MustCompile(LINK_GENERATOR_REGEX)
	}
	return r_LinkGenerator
}
