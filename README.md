# holamundo
 the lambda funciontion holamundo is a test concept to check the effor that is necessary to apply to get a lmabda funcion
 done and querying a database as dynamoDB

## Deploy Go Lambda functions with .zip file archives
the following steps need to be follow to get the zip file done
inside the root directory in wich are the go code, execute following commands
```elixir
 1. go get github.com/aws/aws-lambda-go/lambda
 2. go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest
 3. set GOOS=linux
4. set GOARCH=amd64
5. set CGO_ENABLED=0
6. go build -tags lambda.norpc -o bootstrap main.go
7. %USERPROFILE%\Go\bin\build-lambda-zip.exe -o myFunction.zip bootstrap
```
   

