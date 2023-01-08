go build -o build/sub.exe ./cmd/sub/sub.go
go build -o build/sub_csv.exe ./cmd/sub/sub_csv.go
go build -o build/pub.exe ./cmd/pub/pub.go
go build -o build/api.exe ./cmd/api/api.go

mkdir .\build\Donnees\

COPY .\config.yaml .\build\config.yaml
xcopy /E .\cmd\api\ .\build\ /EXCLUDE:exclude.txt