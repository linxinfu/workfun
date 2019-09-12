default: build-mac

build-mac:
	go build -o bin/workfun ./cmd/workfun

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o bin/workfun.exe ./cmd/workfun
