package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Queue[T any] struct {
	client *sqs.Client
	queue  string
}

func NewQueue[T any](queue string, options aws.Config) *Queue[T] {
	client := sqs.NewFromConfig(options)

	return &Queue[T]{
		queue:  queue,
		client: client,
	}
}

func (q *Queue[T]) Send(data T) error {
	parsed, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal message, %v", err)
	}

	message := string(parsed)

	input := &sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &q.queue,
	}

	_, err = q.client.SendMessage(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send message, %v", err)
	}

	return nil
}

func (q *Queue[T]) Receive() ([]T, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl: &q.queue,
	}

	output, err := q.client.ReceiveMessage(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to receive message, %v", err)
	}

	var messages []T
	for _, message := range output.Messages {
		var m T
		err := json.Unmarshal([]byte(*message.Body), &m)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal message, %v", err)
		}
		messages = append(messages, m)
		defer func() {
			_, err := q.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
				QueueUrl:      &q.queue,
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				fmt.Println("failed to delete message", err)
			}
		}()
	}

	return messages, nil
}
