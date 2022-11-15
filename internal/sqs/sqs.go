package sqs

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

var timeout = 30

func writeToFile(writer *os.File, line string) {
	_, err := writer.WriteString(line + "\n")
	if err != nil {
		fmt.Println("Failed to write to file", err)
	}
}

func ReceiveMessageFromQueue(cfg aws.Config, c context.Context, queueName string, outputFile string) {
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
	fmt.Println("Retriving message from", queueURL)

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Fail to create output file", outputFile, err)
	}

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
			totalMessage += len(msgResult.Messages)
			fmt.Println("Retrieved", len(msgResult.Messages), "messages, total:", totalMessage)

			for _, msg := range msgResult.Messages {
				fmt.Println("Message ID:     " + *msg.MessageId)
				fmt.Println("Message Handle: " + *msg.ReceiptHandle)
				fmt.Println("Message Body: " + *msg.Body)
				writeToFile(file, *msg.Body)
			}
		} else {
			fmt.Println("No messages found")
			fmt.Println("Retrieve %d messages in total.", totalMessage)
			return
		}
	}
}
