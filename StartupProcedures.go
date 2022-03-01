package main

import (
	"fmt"
	"github.com/Arion-Kun/GoLaunch"
	"strings"
)

func initializeStartup() {
	lookupConfig()

	if GoLaunch.Contains("-F") {
		getFile("-F")
	} else if GoLaunch.Contains("--file") {
		getFile("--file")
	}

	if GoLaunch.Contains("-L") {
		getLinks("-L")
	} else if GoLaunch.Contains("--link") {
		getLinks("--link")
	}
	if markDirty {
		StartWorker()
	}

	if !markDirty {
		//If Silent Mode is NOT on print.
		if !Silent {
			println("No file or link found, please run the program without any arguments or links to see the help screen.")
		}
	}
}

func containsArgs(args ...string) bool {
	for i := range args {
		if GoLaunch.Contains(args[i]) {
			return true
		}
	}
	return false
}
func lookupConfig() {
	if containsArgs("-h", "--help") {
		printHelp()
	}
	if containsArgs("-S", "--silent") {
		Silent = true
	}
	if containsArgs("-A", "--async") {
		Async = true
		if containsArgs("-Sr", "--sort") {
			Sort = true
		}
	}
	if containsArgs("-D", "--download") && containsArgs("-O", "--output") && !Async {
		download = true
		contains1, val1 := GoLaunch.TryGetValue("-O")
		if contains1 {
			cachedFolderLocation = val1[0]
		} else {
			contains2, val2 := GoLaunch.TryGetValue("--output")
			if contains2 {
				cachedFolderLocation = val2[0]
			} else {
				download = false
			}
		}
	}
}

var markDirty bool
var Silent bool
var Async bool
var Sort bool
var download bool
var cachedFolderLocation string

func padPrintln(prefix string, suffix string) {
	println(strings.Repeat(" ", 4), prefix, strings.Repeat(" ", clamp(14-len(prefix))), suffix)
}

var clamp = func(in int) int {
	if in < 0 {
		return 0
	}
	return in
}

func padPrintInfo(str string) {
	println(strings.Repeat(" ", 22), str)
}

const blue = "\033[1;34m"
const reset = "\033[0m"

func printHelp() {
	markDirty = true

	println("GoZippy v1.0 - github.com/Arion-Kun/GoZippy")
	println("A simple program to directly download ZippyShare links.\n")
	println("Program Commands / Launch Arguments:\n")
	padPrintln("-h, --help", "| Prints this help screen.")
	padPrintln("-A, --async", "| Checks multiple links at once instead of one at a time.")
	padPrintln("-D, --download", "| Downloads the link(s) specified. (Only works if output is used. Async is unsupported.)") // TODO: Implement
	padPrintln("-F, --file", "| Reads all links from the file.")
	padPrintInfo("(Example: " + blue + "./GoZippy.exe -F 'C:\\Users\\User\\Desktop\\ZippyLinks.txt'" + reset + ")")
	padPrintln("-L, --link", "| Downloads the link(s) specified.")
	padPrintInfo("(Example: " + blue + "./GoZippy.exe -L 'https://www3.zippyshare.com/v/CDCi2wVT/file.html'" + reset + " )")
	padPrintInfo("(MultiLink Example: " + blue + "./GoZippy.exe -L 'https://www20.zippyshare.com/v/oRFjDgWy/file.html' 'https://www20.zippyshare.com/v/GTU4Fiku/file.html'" + reset + " )")
	padPrintln("-S, --silent", "| Suppresses all output except direct links to std-out.")
	padPrintln("-Sr, --sort", "| Outputs the links in the same order it was found in the file. (Only works if async and output is used.)")
	println("\nTypical Example: " + blue + "./GoZippy.exe -L 'https://www3.zippyshare.com/v/CDCi2wVT/file.html'" + reset)
	println("Linux Example: From file, async + silent + sort to std-out/file: " + blue + "\n./GoZippy -F links.txt -A -S -Sr >> output.txt" + reset)
	println(fmt.Sprintf("Download To Folder Example:%s .\\GoZippy.exe -F 'C:\\Users\\User\\Desktop\\links.txt' -D -O 'C:\\Users\\User\\Desktop\\output\\'%s", blue, reset))
}
