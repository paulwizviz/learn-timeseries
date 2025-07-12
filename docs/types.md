# Types

* [By Implementation](#by-implementation)
* [InfluxDB vs Prometheus](#influxdb-vs-prometheus)

## By Implementation

* Purpose-Built
  * InfluxDB
  * Prometheus
  * OpenTSDB
  * Graphite
  * VictoriaMetrics
  * QuestDB
* On Relational DB
  * TimescaleDB (PostgreSQL extension)
  * ClickHouse (columnar DB with time series optimisations)
  * MySQL with custom schema/indexing
* On NoSQL
  * Cassandra (used in KairosDB)
  * ScyllaDB
  * Apache HBase
* Embedded
  * TDengine
  * RRDTool
  * Druid (hybrid analytics)
* Streaming and Even-Based System
  * Apache Kafka with ksqlDB
  * Apache Druid
  * Apache Pinot

Summary of differences

| Type                          | Key Examples                       | Strengths                                      |
|-------------------------------|------------------------------------|------------------------------------------------|
| Purpose-Built TSDBs           | InfluxDB, Prometheus, QuestDB      | High ingestion rates, time-based queries       |
| RDBMS Extensions              | TimescaleDB, ClickHouse            | SQL support, ACID properties                   |
| NoSQL / Wide-Column DBs       | Cassandra, HBase, ScyllaDB         | Horizontal scaling, high availability (HA)     |
| Embedded / Lightweight        | TDengine, RRDTool, VictoriaMetrics | Low resource usage, edge and IoT use cases     |
| Streaming/Event Systems       | Kafka + ksqlDB, Druid, Pinot       | Real-time analytics, streaming data support    |

## InfluxDB vs Prometheus

* Use Case Focus
  * InfluxDB is built for generic time series data: good for sensors, logs, IoT, financial ticks, etc.
  * Prometheus is designed primarily for monitoring systems, like server metrics, application performance, etc.
* Data Ingestion
  * InfluxDB supports push-based ingestion, using REST API, Telegraf, or direct client libraries.
  * Prometheus pulls data via scraping endpoints, making it great for ephemeral infrastructure in Kubernetes.
* Query Language
  * Prometheus uses PromQL, powerful for time-based functions, aggregations, and alerts.
  * InfluxDB v1 used InfluxQL (SQL-like); v2+ uses Flux (functional); v3 (2023+) now supports SQL on top of Apache Arrow.
* Retention and Downsampling
  * InfluxDB natively supports retention policies, downsampling, and continuous queries.
  * Prometheus relies on retention flags and external storage systems (e.g., Thanos, Cortex) for long-term or large-scale storage.
* Scalability
  * Prometheus is a single-node by design (HA setups require external systems like Thanos or Cortex).
  * InfluxDB v3 is cloud-native and horizontally scalable using object storage and decoupled compute.
* Cardinality Handling
  * InfluxDB v3 improves handling of high-cardinality workloads.
  * Prometheus struggles with high-cardinality metrics, which can impact performance or cause out-of-memory errors.

Differences by features

| Feature / Aspect       | InfluxDB                                 | Prometheus                                   |
|------------------------|------------------------------------------|----------------------------------------------|
| Primary Use Case       | General-purpose time series DB           | Metrics and monitoring for systems/apps      |
| Data Model             | Measurement + tags + fields + time       | Metric name + labels + timestamp             |
| Storage Engine         | TSM (v1), IOx (v2/3, columnar)           | Custom time-series storage engine            |
| Query Language         | InfluxQL (v1), Flux (v2+), SQL (v3+)     | PromQL                                       |
| Push vs Pull           | Push-based ingestion                     | Pull-based scraping                          |
| Retention Policies     | Native support with configurable policies| Handled via scrape interval and retention    |
| Alerting               | Kapacitor or Flux scripts                | Built-in Alertmanager                        |
| Long-Term Storage      | Native (e.g., object store in v3)        | Remote Write (Thanos, Cortex, Mimir, etc.)   |
| Scalability            | Horizontal (v3), vertical (v1/v2)        | Single-node; external systems for scaling    |
| Metric Cardinality     | Handles high cardinality well (v3)       | Struggles with high cardinality              |
| Best For               | IoT, industrial, financial, sensor data  | Infrastructure monitoring, DevOps metrics    |
| Deployment             | OSS (v1), Cloud (v2/v3), serverless v3   | OSS, Helm charts, Kubernetes Operator        |

Use case scenarios

| Scenario                             | Recommended Tool                   |
|--------------------------------------|------------------------------------|
| System / application monitoring      | Prometheus                         |
| Cloud-native metrics with dashboards | Prometheus                         |
| Industrial IoT sensor data           | InfluxDB                           |
| Large cardinality financial ticks    | InfluxDB v3                        |
| Long-term storage in object stores   | InfluxDB v3 or Prometheus + Thanos |
| SQL-based analytics on time series   | InfluxDB v3                        |
