export GOOS="windows"
export GOARCH="amd64"
go build -o "./bin/" -ldflags "-s -w" ..
#Reset to defaults
export GOOS=''
export GOARCH=''
echo 'Windows(amd64) Finished.'