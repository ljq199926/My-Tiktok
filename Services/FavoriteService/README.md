# FavoriteService Service

This is the FavoriteService service

Generated with

```
micro new --namespace=go.micro --type=service FavoriteService
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.FavoriteService
- Type: service
- Alias: FavoriteService

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./FavoriteService-service
```

Build a docker image
```
make docker
```