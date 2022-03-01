$env:GOOS = "linux"
$env:GOARCH="amd64"
go build -o "./bin/" -ldflags "-s -w" ..
#Reset to defaults
$env:GOOS=''
$env:GOARCH=''
Write-Output 'Linux(amd64) Finished.'
