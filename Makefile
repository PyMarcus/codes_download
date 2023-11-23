run: build
	@./bin/codes_download perl true

build:
	@go build -o bin/codes_download 
