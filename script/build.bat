@echo off
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go env
go build -a -v -o ./dist/webhook-linux webhook/src
SET GOOS=windows
go env
go build -v -o ./dist/webhook-win.exe webhook/src