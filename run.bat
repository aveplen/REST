@echo off
del cmd.exe
go build -v ./cmd/
cmd.exe
exit 0