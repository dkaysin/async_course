package event_writer

import (
	"async_course/auth"
	schema "async_course/schema_registry"

	"log/slog"
	"os"
)

type EventWriter struct {
	TopicWriterAccount *TopicWriter
	SchemaRegistry     *schema.SchemaRegistry
}

func NewEventWriter(brokers []string, sr *schema.SchemaRegistry) *EventWriter {
	return &EventWriter{
		TopicWriterAccount: newTopicWriter(brokers, auth.KafkaTopicAccount),
		SchemaRegistry:     sr,
	}
}

func (er *EventWriter) Close() {
	if err := er.TopicWriterAccount.w.Close(); err != nil {
		slog.Error("failed to close writer", "error", err)
		os.Exit(1)
	}
}
