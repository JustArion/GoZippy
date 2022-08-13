export GOOS="linux"
export GOARCH="arm64"
go build -o "./bin/GoZippy_Linux_ARM64" -ldflags "-s -w" ..
#Reset to defaults
export GOOS=''
export GOARCH=''
echo 'Linux(arm64) Finished.'