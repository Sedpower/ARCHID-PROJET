go clean ./...
go install ./...
cd cmd/api
%GOPATH%/bin/swag.exe init -g api.go
cd ../..
%GOPATH%/bin/api.exe
