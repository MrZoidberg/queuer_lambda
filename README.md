# Queuer

Simple AWS Lambda that accepts a proxied call from API Gateway and stores the request in SQS. This is useful when you need to process the request asynchronously and return a response to the client as soon as possible. The application is written in Go and uses the AWS SDK for Go.

## Environment Variables

The application uses the following environment variables:

- `AWS_REGION`: The AWS region where the SQS queue is located (by default it will use the same region as the Lambda).
- `QUEUE_URL`: The URL of the SQS queue where the requests will be sent.

## Logging

The application uses the `zap` logging library. If the `DEBUG` environment variable is set to `true`, it will use a development logger with full request logging and also it will add the full request to the response (echo function).

## Building

To build the application, run the following command:

```bash
make build
```

It will produce two zip files in the .bin directory: `function-amd64.zip` and `function-arm64.zip`. The first one is for x86_64 architectures and the second one is for ARM64 architectures.
