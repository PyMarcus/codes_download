run: build
	@./bin/codes_download C true

build:
	@go build -o bin/codes_download 
