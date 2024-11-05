run: build
	bin/server
build:
	go build -o bin/server server/server.go
	env GOOS=js GOARCH=wasm go build -o static/main.wasm main.go
