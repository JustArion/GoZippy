<a href="https://github.com/Arion-Kun/GoZippy/blob/main/go.mod">
	<img height=28 src="https://img.shields.io/github/go-mod/go-version/Arion-Kun/GoZippy?style=flat-square">
</a>
<a href="https://github.com/Arion-Kun/GoZippy/blob/main/.github/workflows/RunTests.yml">
  <img height=28 src="https://img.shields.io/github/workflow/status/Arion-Kun/GoZippy/Function%20Tests?label=Tests&style=for-the-badge">
</a>

#### ⚙️Requirements:
- [GoLang](https://golang.org/dl/)

### Build:
```
git clone https://github.com/Arion-Kun/GoZippy
cd BuildScripts
./Build_Windows_AMD.ps1
```

Alternative:
```sh
git clone https://github.com/Arion-Kun/GoZippy
go build .
```
Further alternatives include selecting the relevant architecture from the BuildScripts folder.

## Launch Arguments:
>-h, --help | Prints this help screen.  
-A, --async | Checks multiple links at once instead of one at a time.  
-D, --download | Downloads the link(s) specified. (Async is unsupported.)  
-F, --file | Reads all links from the file.  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Example: ./GoZippy.exe -F 'C:\Users\User\Desktop\ZippyLinks.txt')  
-L, --link | Downloads the link(s) specified.  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(Example: ./GoZippy.exe -L 'https://www3.zippyshare.com/v/CDCi2wVT/file.html' )  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;(MultiLink Example: ./GoZippy.exe -L 'https://www20.zippyshare.com/v/oRFjDgWy/file.html' 'https://www20.zippyshare.com/v/GTU4Fiku/file.html' )  
-S, --silent | Suppresses all output except direct links to std-out.  
-Sr, --sort | Outputs the links in the same order it was found in the file. (Only works if async and output is used.)

### 📝Simple Link -> Direct Link:
```
./GoZippy.exe -L 'https://www3.zippyshare.com/v/CDCi2wVT/file.html'
```
Output:
```
https://www3.zippyshare.com/d/CDCi2wVT/49267/Gillette%20%2c%20the%20best%20a%20man%20can%20get.wav
```

### 📝Linux Example:

```
/GoZippy -F links.txt -A -S -Sr >> output.txt
```

### 📝Download To Folder:
```
.\GoZippy.exe -F 'C:\Users\User\Desktop\links.txt' -D 'C:\Users\User\Desktop\output\'
```

#### Credits:

[Scotow - Inspiration / Rust Alternative](https://github.com/scotow/zippyst)
<br>
[schollz - ProgressBar](https://github.com/schollz/progressbar/)
<br>
[Knetic - Expression Evaluation](https://pkg.go.dev/github.com/Knetic/govaluate)
