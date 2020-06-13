# Elrond Exporter

Prometheus exporter for Elrond node  
:warning: This is PoC, don't use in production

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

You can now scrape metrics from http://localhost:8888/metrics

### About

Exporter is inspired of [zabbix-elrond-plugin](https://github.com/arcsoft-ro/zabbix-elrond-plugin)