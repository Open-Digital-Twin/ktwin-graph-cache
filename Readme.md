# KTWIN Graph Store

KTWIN Graph Store service.

## Build

```sh
docker build -t ghcr.io/open-digital-twin/ktwin-graph-store:0.1 .
```

## Push Container

```sh
docker push ghcr.io/open-digital-twin/ktwin-graph-store:0.1
```

## Load in Kind Development Environment

```sh
docker build -t dev.local/open-digital-twin/ktwin-graph-store:0.1 .
kind load docker-image dev.local/open-digital-twin/ktwin-graph-store:0.1
```

## CURL Commands

```sh
kubectl run curl \
    --image=curlimages/curl --rm=true --restart=Never -ti -- \
    -X GET -v \
    -H "content-type: application/json"  \
    http://ktwin-graph-store.ktwin.svc.cluster.local/api/v1/twin-graph
```

```sh
kubectl run curl \
    --image=curlimages/curl --rm=true --restart=Never -ti -- \
    -X GET -v \
    -H "content-type: application/json"  \
    http://ktwin-graph-store.ktwin.svc.cluster.local/api/v1/twin-graph/ngsi-ld-city-device
```

```sh
kubectl run curl \
    --image=curlimages/curl --rm=true --restart=Never -ti -- \
    -X GET -v \
    -H "content-type: application/json"  \
    http://ktwin-graph-store.ktwin.svc.cluster.local/health
```
