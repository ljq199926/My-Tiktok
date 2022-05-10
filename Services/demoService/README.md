# DemoService Service

This is the DemoService service

Generated with

```
micro new --namespace=go.micro --type=service demoService
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.demoService
- Type: service
- Alias: demoService

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
./demoService-service
```

Build a docker image
```
make docker
```