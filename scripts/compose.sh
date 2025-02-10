#!/bin/bash

export NETWORK_NAME="local_network"
export INFLUX_IMAGE=influxdb:2.6

COMMAND=$1

case $COMMAND in
    "clean")
        docker rmi -f ${INFLUX_IMAGE}
        docker rmi -f $(docker images --filter "dangling=true" -q)
        rm -rf ./deployments/influxdb_data
        ;;
    "start")
        docker compose -f ./deployments/docker-compose.yaml up
        ;;
    "stop")
        docker compose -f ./deployments/docker-compose.yaml down
        ;;
    *)
        echo "Usage: $0 [start | stop]"
        ;;
esac