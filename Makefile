run: build
	@./bin/codes_download python true 2023

build:
	@go build -o bin/codes_download 
