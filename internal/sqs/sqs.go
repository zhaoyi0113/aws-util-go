package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

var timeout = 30

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
	fmt.Println("Retriving message from %s", queueURL)
	totalMessage := 0
	for {
		gMInput := &sqs.ReceiveMessageInput{
			MessageAttributeNames: []string{
				string(types.QueueAttributeNameAll),
			},
			QueueUrl:            queueURL,
			MaxNumberOfMessages: 10,
			VisibilityTimeout:   int32(timeout),
		}
		msgResult, err := client.ReceiveMessage(c, gMInput)
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			fmt.Println(err)
			return
		}
		if msgResult.Messages != nil {
			fmt.Println("Retrieved %d messages, total %d", len(msgResult.Messages), totalMessage)
			totalMessage += len(msgResult.Messages)

			for _, msg := range msgResult.Messages {
				fmt.Println("Message ID:     " + *msg.MessageId)
				fmt.Println("Message Handle: " + *msg.ReceiptHandle)
				fmt.Println("Message Body: " + *msg.Body)
			}
		} else {
			fmt.Println("No messages found")
			fmt.Println("Retrieve %d messages in total.", totalMessage)
			return
		}
	}
}

func getQueueURL(c context.Context, client *sqs.Client, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return client.GetQueueUrl(c, input)
}
