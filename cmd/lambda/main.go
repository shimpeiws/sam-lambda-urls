package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/labstack/gommon/log"

	"github.com/shimpeiws/sam-lambda-urls/http_handlers"
)

func main() {
	handler, err := http_handlers.DefaultHandler()
	if err != nil {
		log.Error("aaa!!!")
		os.Exit(1)
	}

	lambda.Start(httpadapter.NewV2(handler).ProxyWithContext)
}
