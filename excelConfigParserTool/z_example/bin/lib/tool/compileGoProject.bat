echo goarch = %1
echo goos = %2
echo gopath = %3
echo output = %4
echo input = %5

set GOARCH=%1
set GOOS=%2
set GOPATH=%3

go build -o %4 %5