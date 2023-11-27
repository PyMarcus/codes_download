run: build
	@./bin/codes_download cobol true

build:
	@go build -o bin/codes_download 
