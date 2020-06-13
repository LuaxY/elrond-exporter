# Elrond Exporter

Prometheus exporter for Elrond node

### Metrics

All metrics are prefixed by `erd_`

### Build

```shell script
go build ./cmd/elrond-exporter
```

### Run

```shell script
./elrond-export --port 8888 --interval 5
```