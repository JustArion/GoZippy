$env:GOOS = "linux"
$env:GOARCH="arm64"
go build -o "./bin/GoZippy_Linux_ARM64" -ldflags "-s -w" ..
#Reset to defaults
$env:GOOS=''
$env:GOARCH=''
Write-Output 'Linux(arm64) Finished.'
