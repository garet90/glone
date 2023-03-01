all: !tests

!tests:
	GOOS=js GOARCH=wasm go build -o ./bin/test/webgl.wasm ./test/webgl