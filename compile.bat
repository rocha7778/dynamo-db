del bootstrap
del main.zip

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
 %USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip bootstrap