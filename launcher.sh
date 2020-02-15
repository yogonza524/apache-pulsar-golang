#!/bin/bash

PULSAR_POD=$(podman pod ps | grep pulsar)

if [[ -z $PULSAR_POD ]]; then
    echo "[Pod] Creating..."
    podman pod create --name=pulsar --publish 6650:6650
fi

echo "[Pod] Created..."

PULSAR_CONTAINER=$(podman ps | grep pulsar-container)

if [[ -z $PULSAR_CONTAINER ]]; then
    echo "[Container] Start..."
    podman run -dt --pod=pulsar --name=pulsar-container apachepulsar/pulsar bin/pulsar standalone
fi

echo "[Container] Running.."