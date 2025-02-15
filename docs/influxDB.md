# InfluxDB

This covers aspects of InfluxDB:

* [Deployment](#deployment)
* [Working Examples](#working-examples)
* [References](#references)

## Key concepts

* Bucket == database
* Measurements == table
* Tags & Fields converted to columns.
    * Tags store metadata
    * Fields actual values

## Deployments

* [Docker compose](#docker-compose)

### Docker compose

Minimum environment variables for Docker container setting.

* **DOCKER_INFLUXDB_INIT_MODE**
* **DOCKER_INFLUXDB_INIT_ORG**
* **DOCKER_INFLUX_BUCKET**
* **DOCKER_INFLUXDB_INIT_ADMIN_TOKEN**

The following are potential values for **DOCKER_INFLUXDB_INIT_MODE**

| Value | Description |
|---|---|
| **setup** | The recommended value. Automatically creates an initial admin user, organisation, bucket, and authentication token on first startup. |
| **upgrade** |	Used when upgrading an existing InfluxDB instance. Keeps existing data while upgrading metadata. |

To generate tokens and assign to variable **DOCKER_INFLUXDB_INIT_ADMIN_TOKEN**, these are the following options:

* **Option 1**: Using openssl. Run the command `openssl rand -base64 32`

## Working Examples

* V2
    * [Example 1](../examples/influx/v2/ex1/main.go) - This example demonstrates writing to influxDB using point types and write blocking operations.
    * [Example 2](../examples/influx/v2/ex2/main.go) - This example demonstrates writing to influxDB using point types and async write operations. 

## References

* [Official Documentation](https://docs.influxdata.com/)
* [Flux QL](https://docs.influxdata.com/influxdb/cloud/reference/syntax/flux/flux-vs-influxql/)
