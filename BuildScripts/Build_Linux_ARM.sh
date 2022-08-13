export GOOS="linux"
export GOARCH="arm"
go build -o "./bin/GoZippy_Linux_ARM" -ldflags "-s -w" ..
#Reset to defaults
export GOOS=''
export GOARCH=''
echo 'Linux(arm) Finished.'