cd ../\
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go env
go build -a -v -o ./dist/webhook-windows.exe webhook/src
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go env
go build -v -o ./dist/webhook-linux webhook/src