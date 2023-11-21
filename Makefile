run: build
	@./bin/codes_download python true

build:
	@go build -o bin/codes_download 
