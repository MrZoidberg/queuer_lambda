package main

type options struct {
	SQS struct {
		QueueURL  string `env:"SQS_QUEUE_URL" short:"q" long:"queue_url" required:"true" description:"The URL of the SQS queue to send messages to"`
		AWSRegion string `env:"AWS_REGION" short:"r" long:"aws_region" required:"true" description:"The AWS region the SQS queue is in"`
	}
	Debug      bool   `env:"DEBUG" short:"d" long:"debug" description:"Enable debug logging"`
	OKResponse string `env:"OK_RESPONSE" short:"o" long:"ok_response" description:"The response body to return when the message is sent successfully"`
}

func NewOptions() *options {
	a := options{}
	return &a
}
