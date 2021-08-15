govvv build -pkg github.com/psidex/CrowsNest/cmd -o crowsnest-windows-amd64.exe .\main.go

$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; $Env:CGO_ENABLED = 0;

govvv build -pkg github.com/psidex/CrowsNest/cmd -o ./crowsnest-linux-amd64 ./main.go
