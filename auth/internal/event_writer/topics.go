package event_writer

import "async_course/auth"

type EventWriter struct {
	TopicWriterAccount *TopicWriter
}

func NewEventWriter(brokers []string) *EventWriter {
	return &EventWriter{
		TopicWriterAccount: newTopicWriter(brokers, auth.KafkaTopicAccount),
	}
}
