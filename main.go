package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	topicArn = "arn:aws:sns:us-east-1:000000000000:meu-topico"
)

func main() {
	sessao := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1"),
	}))

	svc := sns.New(sessao)

	publishParams := &sns.PublishInput{
		Message:  aws.String("Mensagem teste"),
		TopicArn: aws.String(topicArn),
	}

	_, err := svc.Publish(publishParams)
	if err != nil {
		println(err)
		return
	}

	println("Mensagem enviada com sucesso")
}
