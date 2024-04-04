package main

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v10"
)

var version = "unknown"
var commit = "unknown"
var date = "unknown"

var logger *zap.Logger

func main() {
	fmt.Printf("Query version: %s, commit: %s, date: %s\n", version, commit, date)

	opts := NewOptions()
	if err := env.Parse(opts); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", opts)

	if opts.Debug {
		fmt.Println(os.Environ())
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}

	if logger == nil {
		log.Fatalf("Cannot initialize zap logger: %v", logger)
	}

	h := &handler{opts: opts}
	lambda.Start(h.handleRequest)
}
