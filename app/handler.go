package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"
)

type handler struct {
	opts *options
}

func (h *handler) handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if ctx == nil {
		logger.Error("context is nil")
	}
	if h.opts == nil {
		logger.Error("opts is nil")
	}
	if h.opts.SQS.AWSRegion == "" {
		logger.Error("opts.SQS.AWSRegion is empty")
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(h.opts.SQS.AWSRegion))
	if err != nil {
		// logger.Error("unable to load SDK config", zap.Error(err))
		return events.APIGatewayV2HTTPResponse{Body: "Initialization failed: " + err.Error(), StatusCode: 500}, nil
	}

	client := sqs.NewFromConfig(cfg)

	req, err := json.Marshal(request)
	if err != nil {
		logger.Error("unable to marshal request", zap.Error(err))
		json_err, _ := json.Marshal(err)
		return events.APIGatewayV2HTTPResponse{Body: string(json_err), StatusCode: 500, Headers: map[string]string{"Content-Type": "application/json"}}, nil
	}

	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(string(req)),
		QueueUrl:    &h.opts.SQS.QueueURL,
	})

	if err != nil {
		logger.Error("unable to send message", zap.Error(err))
		json_err, _ := json.Marshal(err)
		return events.APIGatewayV2HTTPResponse{Body: string(json_err), StatusCode: 500, Headers: map[string]string{"Content-Type": "application/json"}}, nil
	}

	if h.opts.Debug {
		logger.Debug("message sent", zap.String("body", request.Body))
	}

	if h.opts.OKResponse != "" {
		return events.APIGatewayV2HTTPResponse{Body: h.opts.OKResponse, StatusCode: 200, Headers: map[string]string{"Content-Type": "application/json"}}, nil
	}

	body := map[string]interface{}{
		"status": "ok",
	}
	if h.opts.Debug {
		body["request"] = request
	}
	body_json, _ := json.Marshal(body)
	return events.APIGatewayV2HTTPResponse{Body: string(body_json), StatusCode: 200, Headers: map[string]string{"Content-Type": "application/json"}}, nil
}
