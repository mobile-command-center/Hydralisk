.PHONY: build clean deploy all
lambda:
	@GOOS=linux go build -o main main.go

compression:
	@zip main.zip main

build: lambda compression

deploy:
	@aws lambda update-function-code --function-name hydralisk-dev-registration --zip-file fileb://main.zip

all: build deploy

clean:
	@rm main
	@rm main.zip
