$env:GOOS = "windows"
$env:GOARCH="amd64"
go build -o "./bin/" -ldflags "-s -w" ..
#Reset to defaults
$env:GOOS=''
$env:GOARCH=''
Write-Output 'Build to Windows(amd64) Complete.'
