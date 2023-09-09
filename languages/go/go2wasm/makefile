name := hello.wasm
exec := $(shell go env GOROOT)/misc/wasm/go_js_wasm_exec

run-wasm: hello.wasm
	GOOS=js GOARCH=wasm go run -exec=$(exec) .

hello.wasm: go.mod
	GOOS=js GOARCH=wasm go build .

clean:
	rm *.wasm

browser: index.html hello.wasm wasm_exec.js

wasm_exec.js:
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

go.mod:
	go mod init $(name)
