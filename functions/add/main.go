package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	A int `json:"VAR_A"`
	B int `json:"VAR_B"`
}

func HandleRequest(ctx context.Context, e *MyEvent) (*int, error) {
	fmt.Println(e)
	if e == nil {
		return nil, fmt.Errorf("received nil event %v", e)
	}

	res := add(e.A, e.B)
	return &res, nil
}

func add(a, b int) int {
	return a + b
}

func main() {
	lambda.Start(HandleRequest)
}
