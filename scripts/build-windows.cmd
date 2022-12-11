set GOOS=windows
set GOARCH=amd64
go build -o ../bin/prime-numbers-tester.exe ../cmd/prime-numbers-tester.go
pause