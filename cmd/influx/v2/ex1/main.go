package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Event struct {
	TS    time.Time `json:"time_stamp"`
	Value int       `json:"value"`
}

func generateEvent(ctx context.Context) chan Event {
	event := make(chan Event, 1)
	go func() {
		defer close(event)
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				event <- Event{
					TS:    time.Now(),
					Value: rand.Intn(10) + 1,
				}
			}
		}
	}()
	return event
}

func writeToInfluxDB(ctx context.Context, events chan Event, cred, org, bucket string) {
	// Connect to InfluxDB
	influxClient := influxdb2.NewClient("http://localhost:8086", cred)
	defer influxClient.Close()

	writeAPI := influxClient.WriteAPIBlocking(org, bucket)

	for event := range events {
		point := influxdb2.NewPoint(
			"events",
			map[string]string{},
			map[string]interface{}{
				"timestamp": event.TS,
				"value":     event.Value,
			},
			time.Now(),
		)
		if err := writeAPI.WritePoint(ctx, point); err != nil {
			log.Printf("Error writing to InfluxDB: %v", err)
		} else {
			log.Printf("Event written to InfluxDB successfully: %+v", event)
		}
	}
}

func main() {

	// Get InfluxDB credentials from environment variables
	influxCred := os.Getenv("INFLUXDB_TOKEN")
	if influxCred == "" {
		log.Fatal("InfluxDB token is not specified in INFLUXDB_TOKEN")
	}

	influxBucket := os.Getenv("DOCKER_INFLUXDB_INIT_BUCKET")
	if influxBucket == "" {
		log.Fatal("InfluxDB token is not specified in DOCKER_INFLUXDB_INIT_BUCKET")
	}

	influxOrg := os.Getenv("DOCKER_INFLUXDB_INIT_ORG")
	if influxOrg == "" {
		log.Fatal("InfluxDB token is not specified in DOCKER_INFLUXDB_INIT_ORG")
	}

	ctx, cancel := context.WithCancel(context.Background())
	events := generateEvent(ctx)
	writeToInfluxDB(ctx, events, influxCred, influxOrg, influxBucket)

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt, syscall.SIGTERM)
	<-osSig
	cancel()
}
