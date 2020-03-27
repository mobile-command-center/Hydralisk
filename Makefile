lambda:
	@GOOS=linux go build -o cmd/lambda/main cmd/lambda/main.go

compression:
	@zip cmd/lambda/main.zip cmd/lambda/main

build: lambda compression

deploy:
	@aws lambda update-function-code --function-name hydralisk-dev-registration --zip-file fileb://cmd/lambda/main.zip

all: build

clean:
	@rm cmd/lambda/main
	@rm cmd/lambda/main.zip