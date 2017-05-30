json-to-statsd
======================

Sends info from a http (GET) service as gauge to statsd every 'x' seconds

# Help
```sh
$ json-to-statsd -h
```

# Map
[gjson](github.com/tidwall/gjson) is used for mapping the metrics from json, for path syntax, check on their docs.

The map format is defined by setting the key for statsd metrics as key and the json path containing the metrics value as value.
```
statsd.gauge.key: json.path
```

## Eg
### json service response
```json
{
    "version": "1.0.0",
    "metrics": {
        "totalHeapBytes": 32212254720,
        "usedHeapBytes": 1292247520,
        "totalThreads": 290
    }
}
```
### map
```
example.heap.total: metrics.totalHeapBytes
example.heap.used: metrics.usedHeapBytes
example.threads.total: metrics.totalThreads
```
