export GOOS='linux'
export GOARCH='amd64'
go build -o "./bin/" -ldflags "-s -w" ..
#Reset to defaults
export GOOS=''
export GOARCH=''
echo 'Linux(amd64) Finished.'