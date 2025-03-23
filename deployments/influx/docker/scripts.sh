#!/bin/bash

if [ "$(basename $(realpath .))" != "learn-timeseries" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export NETWORK_NAME="learn-timeseries_network"
export INFLUX_IMAGE=influxdb:2.6

COMMAND=$1
SUBCOMMAND=$2

function single(){
    local cmd=$1
    case $cmd in
        "clean")
            docker compose -f ./deployments/influx/docker/single/docker-compose.yaml down
            docker rmi -f ${INFLUX_IMAGE}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        "start")
            docker compose -f ./deployments/influx/docker/single/docker-compose.yaml up
            ;;
        "stop")
            docker compose -f ./deployments/influx/docker/single/docker-compose.yaml down
            ;;
        *)
            echo "Usage: $0 single [start | stop]"
            ;;
    esac
}

case $COMMAND in
    "single")
        single $SUBCOMMAND
        ;;
    *)
        echo "Usage: $0 single"
        ;;
esac