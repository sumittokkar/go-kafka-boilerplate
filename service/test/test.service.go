package test

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type TestService struct {
}

func NewTestService() *TestService {
	return &TestService{}
}

func (ps *TestService) Consume(msg *kafka.Message) {
	panic("implement me")
}
