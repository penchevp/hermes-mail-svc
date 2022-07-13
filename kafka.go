package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

func standardKafkaConfig() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers:  []string{"localhost:29092"},
		GroupID:  applicationGroupID,
		MinBytes: 1e2,
		MaxBytes: 10e6,
	}
}

func withTopic(readerConfig kafka.ReaderConfig, topic string) kafka.ReaderConfig {
	readerConfig.Topic = topic
	return readerConfig
}

func withOffset(readerConfig kafka.ReaderConfig, offset int64) kafka.ReaderConfig {
	readerConfig.StartOffset = offset
	return readerConfig
}

func readKafka(ctx context.Context, readerConfig kafka.ReaderConfig, fetchMessageChan chan<- kafka.Message, commitMessageChan <-chan kafka.Message) {
	reader := kafka.NewReader(readerConfig)
	defer reader.Close()

	go func() {
		for {
			m := <-commitMessageChan

			if err := reader.CommitMessages(ctx, m); err != nil {
				log.Err(err)
			}
		}
	}()

	run := true
	for run {
		select {
		case <-ctx.Done():
			{
				run = false
			}
		default:
		}

		m, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Err(err)
			continue
		}

		if m.Value == nil {
			continue
		}

		fetchMessageChan <- m
	}
}
