go build -tags lambda.norpc -o bootstrap main.go
 %USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip bootstrap