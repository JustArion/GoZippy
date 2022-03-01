$env:GOOS = "linux"
$env:GOARCH="arm64"
go build -o "./bin/" -ldflags "-s -w" ..
#Reset to defaults
$env:GOOS=''
$env:GOARCH=''
Write-Output 'Linux(amd64) Finished.'
