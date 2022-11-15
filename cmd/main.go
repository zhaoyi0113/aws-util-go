package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/zhaoyi0113/aws/utils/internal/sqs"
)

func main() {
	c := context.TODO()
	cfg, err := config.LoadDefaultConfig(c, config.WithRegion("ap-southeast-2"))
	if err != nil {
		fmt.Println("Failed to load AWS default config")
		return
	}
	sqs.ReceiveMessageFromQueue(cfg, c, "dev-ams-cqrs-projection-dlq")
}
