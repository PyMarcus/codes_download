run: build
	@./bin/codes_download rust true

build:
	@go build -o bin/codes_download 
