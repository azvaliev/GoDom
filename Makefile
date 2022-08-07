build:
	GOOS=js GOARCH=wasm go build -o ./out.wasm **/*.go 


