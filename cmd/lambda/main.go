package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func registrationHandler(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {
	lambda.Start(registrationHandler)
}
