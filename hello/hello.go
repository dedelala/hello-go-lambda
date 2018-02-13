package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Go handles a Lambda invocation.
func Go() {
	log.Println("Hello log world!")
}

func main() {
	lambda.Start(Go)
}
