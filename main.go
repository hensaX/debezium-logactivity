package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hensaX/debezium-logactivity/model"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	client, _ := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumeTopics("log-activity"),
		kgo.ConsumerGroup("log-activity-service"),
	)
	log.Println("start consume")
	go StartConsumer(ctx, client)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	log.Println("shutdown signal")

	cancel()       // ⬅ stop poll
	client.Close() // ⬅ close kafka connection
}

func StartConsumer(ctx context.Context, client *kgo.Client) {
	for {
		select {
		case <-ctx.Done():
			log.Println("consumer shutdown signal received")
			return
		default:
		}

		fetches := client.PollFetches(ctx)

		// ⛔ WAJIB cek error
		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				// Kalau context cancel → stop loop
				if ctx.Err() != nil {
					log.Println("poll stopped:", err)
					return
				}
				log.Println("poll error:", err)
			}
			continue
		}

		fetches.EachRecord(func(r *kgo.Record) {
			processLogActivity(r)
		})
	}
}

func processLogActivity(r *kgo.Record) {
	msgData := &model.CDCData{}
	err := json.Unmarshal(r.Value, msgData)
	if err != nil {
		fmt.Sprintf("error json unmarshal msg-data %w", err)
	}
	log.Printf(
		"topic=%s partition=%d offset=%d key=%s value=%s",
		r.Topic,
		r.Partition,
		r.Offset,
	)
	log.Printf("\n=============start\n[before] : %v \n[after] : %v\nend=============", msgData.Payload.Before, msgData.Payload.After)

}
