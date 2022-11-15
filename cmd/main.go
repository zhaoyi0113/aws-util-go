package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/zhaoyi0113/aws/utils/internal/sqs"
)

type arguments struct {
	cmd       string
	queueName string
}

func defineFlags() arguments {
	cmd := flag.String("c", "", "AWS command")
	queueName := flag.String("q", "", "SQS queue name")
	flag.Parse()
	args := arguments{cmd: *cmd, queueName: *queueName}
	return args
}

func main() {
	args := defineFlags()
	fmt.Println("cmd", args.cmd, "queueName ", args.queueName)
	if len(args.cmd) == 0 || len(args.queueName) == 0 {
		fmt.Println("Please speicfy command and queue name")
		flag.PrintDefaults()
		return
	}
	c := context.TODO()
	cfg, err := config.LoadDefaultConfig(c, config.WithRegion("ap-southeast-2"))
	if err != nil {
		fmt.Println("Failed to load AWS default config")
		return
	}
	sqs.ReceiveMessageFromQueue(cfg, c, args.queueName)
}
