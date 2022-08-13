$env:GOOS = "linux"
$env:GOARCH="arm"
go build -o "./bin/GoZippy_Linux_ARM" -ldflags "-s -w" ..
#Reset to defaults
$env:GOOS=''
$env:GOARCH=''
Write-Output 'Linux(arm) Finished.'
