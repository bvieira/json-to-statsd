json-to-statsd
======================

Sends info from a http (GET) service as gauge to statsd every 'x' seconds

# Help & Usage
```bash
$ json-to-statsd -h
```

# Map
The map format is defined by setting the key for statsd metrics as key and the json path containing the metrics value as value on a yml file.
```yaml
stats.gauge.key: json.path
```

## JSON path
[gjson](http://github.com/tidwall/gjson) is used for mapping the metrics from json, for path syntax, check on their docs.

## Functions
It's possible to replace values using go native "text/template" syntax, **the template is rendered before applying json values only**

eg. APP_NAME=app1
template:
```yaml
dynamic.metrics.heap.total: metrics.#[name=="{{env "APP_NAME"}}"].totalHeapBytes
```
result:
```yaml
dynamic.metrics.heap.total: metrics.#[name=="app1"].totalHeapBytes
```

## Aditional template functions
### env [key]
get key from env var, returns an error if not exists

### envd [key] [default-value]
get key from env var, returns default-value if not exists

### lower [value]
string to lowercase

### upper [value]
string to uppercase


## Eg.
```bash
$ APP_NAME=app1 json-to-statsd -u http://my-api/status -p example -m map.yml -s 127.0.0.1:8125
```

### json service response
```json
{
    "version": "1.0.0",
    "metrics": [
        {
            "name": "app1",
            "totalHeapBytes": 32212254721,
            "usedHeapBytes": 1292247521,
            "totalThreads": 291
        },
        {
            "name": "app2",
            "totalHeapBytes": 32212254722,
            "usedHeapBytes": 1292247522,
            "totalThreads": 292
        }
    ]
}
```
### map
```yaml
version: version
fixed.app1.name: metrics.0.name
fixed.app2.name: metrics.1.name
dynamic.metrics.heap.total: metrics.#[name=="{{env "APP_NAME"}}"].totalHeapBytes
dynamic.metrics.heap.used: metrics.#[name=="{{env "APP_NAME"}}"].usedHeapBytes
dynamic.metrics.threads.total: metrics.#[name=="{{env "APP_NAME"}}"].totalThreads
```

### result
```yaml
example.version: 1.0.0
example.fixed.app1.name: app1
example.fixed.app2.name: app2
example.dynamic.metrics.heap.total: 32212254721
example.dynamic.metrics.heap.used: 1292247521
example.dynamic.metrics.threads.total: 291
```