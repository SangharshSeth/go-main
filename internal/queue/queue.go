package queue

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSSender struct {
	client   *sqs.Client
	queueURL string
}

func NewSqsSender(queueUrl string, profile string) *SQSSender {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion("ap-south-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &SQSSender{
		client:   sqs.NewFromConfig(cfg),
		queueURL: queueUrl,
	}
}

func (s *SQSSender) SendMessageToSQS(message string) {
	_, err := s.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &s.queueURL,
		MessageBody: &message,
	})
	if err != nil {
        log.Fatalf("failed to send message, %v", err)
    }

    log.Println("Message sent successfully")
}

func SendMessageToSQS(message string) {
	_, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("web-developer"))
	if err != nil {
		log.Fatal(err)
	}
}
