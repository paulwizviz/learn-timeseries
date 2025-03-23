# InfluxDB

* [Concepts](#concepts)
* [Deployment](#deployments)
* [References](#references)

## Concepts

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
| **upgrade** | Used when upgrading an existing InfluxDB instance. Keeps existing data while upgrading metadata. |

To generate tokens and assign to variable **DOCKER_INFLUXDB_INIT_ADMIN_TOKEN**, these are the following options:

* **Option 1**: Using openssl. Run the command `openssl rand -base64 32`

## References

* [Official Documentation](https://docs.influxdata.com/)
* [Flux QL](https://docs.influxdata.com/influxdb/cloud/reference/syntax/flux/flux-vs-influxql/)
* [InfuxDB: Overview, Key Concepts and Demo | Getting Started](https://www.youtube.com/watch?v=gb6AiqCJqP0)
