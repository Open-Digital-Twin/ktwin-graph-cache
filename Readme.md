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

## Get ScyllaDB credentials from secrets

```sh
kubectl get secrets/scylla-auth-token -n scylla --template={{.data}}
```

## Expose ScyllaDB Port

```sh
kubectl port-forward --address 0.0.0.0 -n scylla svc/scylla-client 9042:9042
```

## Run scripts in ScyllaDB cluster

You can manually run the scripts in any of the ScyllaDB nodes.

```sh
kubectl get pods -n ktwin
kubectl exec -ti scylla-us-east-1-us-east-1b-1 bash -n ktwin
cqlsh
```

Now run the CQL commands.
