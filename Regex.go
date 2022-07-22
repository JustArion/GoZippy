package main

import "regexp"

const SCRIPT_REGEX = "(?s)<script type=\"text/javascript\">(.+?)</script>"

/*
	<script type="text/javascript">
	    var a = 38;
	    document.getElementById('dlbutton').omg = "asdasd".substr(0, 3);
	    var b = document.getElementById('dlbutton').omg.length;
	    document.getElementById('dlbutton').href = "/d/PWRXlkLH/"+(Math.pow(a, 3)+b)+"/1mb";
	    if (document.getElementById('fimage')) {
	        document.getElementById('fimage').href = "/i/PWRXlkLH/"+(Math.pow(a, 3)+b)+"/1mb";
	    }
	</script>
*/
//const VARIABLE_REGEX = "var\\s*(\\w+)\\s*=\\s*([\\d+\\-*/%]+);?"
const VARIABLE_REGEX = "var a = (\\d+);[$||\\n]"
const VARIABLE_KEY_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.omg = \\\"([^\\\"]+)\\\""

//const LINK_GENERATOR_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\s*\\+\\s*([\\d\\w\\s+\\-*/%()]+?)\\s*\\+\\s*\"/([/\\w%.-]+)\";?"
//const LINK_GENERATOR_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\s*\\+[^\\\"]+\"/([/\\w%.-]+)\";?"
const LINK_GENERATOR_REGEX = "document\\.getElementById\\('dlbutton'\\)\\.href\\s*=\\s*\"/d/(\\w+)/\"\\+\\(([^\\%]+)\\%1000\\s*\\+[^\\\"]+\"/([/\\w%.-]+)\";?"

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

var r_VariableKey *regexp.Regexp

func GetVariableKeyRegex() *regexp.Regexp {

	if r_VariableKey == nil {
		r_VariableKey = regexp.MustCompile(VARIABLE_KEY_REGEX)
	}
	return r_VariableKey
}

var r_LinkGenerator *regexp.Regexp

func GetLinkGeneratorRegex() *regexp.Regexp {

	if r_LinkGenerator == nil {
		r_LinkGenerator = regexp.MustCompile(LINK_GENERATOR_REGEX)
	}
	return r_LinkGenerator
}
