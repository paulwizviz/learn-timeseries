package main

import (
	"context"
	"go-timescaledb/internal/event"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func writeToInfluxDB(ctx context.Context, events chan event.Weather, cred, org, bucket string) {
	// Connect to InfluxDB
	influxClient := influxdb2.NewClient("http://localhost:8086", cred)
	defer influxClient.Close()

	// Using blocking operations
	writeAPI := influxClient.WriteAPIBlocking(org, bucket)

	for event := range events {
		point := influxdb2.NewPoint(
			"weather",
			map[string]string{ // tags
				"country":  event.Country,
				"location": event.Location,
			},
			map[string]interface{}{ // field
				"temperature": event.Temp,
				"rain":        event.Rain,
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
	events := event.Generate(ctx)
	writeToInfluxDB(ctx, events, influxCred, influxOrg, influxBucket)

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt, syscall.SIGTERM)
	<-osSig
	cancel()
}
