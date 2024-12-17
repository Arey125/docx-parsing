BIN = docx-parsing

all:
	@go build -o bin/$(BIN) cmd/cli/main.go

windows:
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o bin/$(BIN)-win.exe cmd/cli/main.go

run: all
	@./bin/$(BIN) -f ./files/document.docx

help: all
	@./bin/$(BIN) -h
