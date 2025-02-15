package main

import (
	"context"
	"fmt"
	"go-timescaledb/internal/event"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func writeToInfluxDB(events chan event.Weather, cred, org, bucket string) {
	// Connect to InfluxDB
	client := influxdb2.NewClient("http://localhost:8086", cred)
	defer client.Close()

	// Using blocking operations
	writeAPI := client.WriteAPI(org, bucket)
	errors := writeAPI.Errors()
	// Create goroutine to log errors
	go func() {
		for err := range errors {
			fmt.Printf("write error: %s\n", err.Error())
		}
	}()
	defer writeAPI.Flush()

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
		writeAPI.WritePoint(point)
		fmt.Printf("Written point: %v\n", point)
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
	writeToInfluxDB(events, influxCred, influxOrg, influxBucket)

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt, syscall.SIGTERM)
	<-osSig
	cancel()
}
