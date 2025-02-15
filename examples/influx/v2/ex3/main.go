package main

import (
	"context"
	"fmt"
	"go-timescaledb/internal/weather"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func write(ctx context.Context, client influxdb2.Client, events []weather.Event, org, bucket, measurement string) {
	// Using blocking operations
	writeAPI := client.WriteAPIBlocking(org, bucket)
	defer writeAPI.Flush(ctx)

	for _, event := range events {
		point := influxdb2.NewPoint(
			measurement,
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

func query(ctx context.Context, client influxdb2.Client, influxOrg string, influxBucket string, measurement string) error {

	queryAPI := client.QueryAPI(influxOrg)
	ql := fmt.Sprintf(`from(bucket:"%s")
	|> range(start: -1h) 
	|> filter(fn: (r) => r._measurement == "%s")`, influxBucket, measurement)

	log.Println(ql)

	result, err := queryAPI.Query(ctx, ql)
	if err != nil {
		return err
	}
	// check for an error
	if result.Err() != nil {
		return result.Err()
	}
	for result.Next() {
		// Notice when group key has changed
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// Access data
		fmt.Printf("timestamp:%v field: %v value: %v\n", result.Record().Time().Unix(), result.Record().Field(), result.Record().Value())
	}
	return nil
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
	events := weather.GenerateTenEvents()

	// Connect to InfluxDB
	client := influxdb2.NewClient("http://localhost:8086", influxCred)
	defer client.Close()
	measurement := "weather"
	write(context.TODO(), client, events, influxOrg, influxBucket, measurement)
	err := query(context.TODO(), client, influxOrg, influxBucket, measurement)
	if err != nil {
		log.Println("Query Error: ", err)
	}
}
