package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func ReceiveMessageFromQueue(cfg aws.Config, c context.Context, queueName string) {
	client := sqs.NewFromConfig(cfg)
	urlResult, err := client.GetQueueUrl(c, &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}
	queueURL := urlResult.QueueUrl

	gMInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: 1,
		// VisibilityTimeout:   int32(*timeout),
	}
	msgResult, err := client.ReceiveMessage(c, gMInput)
	if err != nil {
		fmt.Println("Got an error receiving messages:")
		fmt.Println(err)
		return
	}
	if msgResult.Messages != nil {
		fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	} else {
		fmt.Println("No messages found")
	}
}

func getQueueURL(c context.Context, client *sqs.Client, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return client.GetQueueUrl(c, input)
}
