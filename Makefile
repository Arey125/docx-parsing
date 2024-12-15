BIN = docx-parsing

all:
	@go build -o bin/$(BIN) cmd/main.go

run: all
	@./bin/$(BIN) -f ./files/document.docx

help: all
	@./bin/$(BIN) -h
