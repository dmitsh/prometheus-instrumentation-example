# Prometheus Instrumentation Example in Go

This project implements a simple web server/client.
The webserver code includes:
 - instrumentation of Prometheus client
 - configuration of custom metrics
 - simulation of status code, duration, and payload size for each HTTP request

The client runs an endless loop of sending HTTP requests to the web server.

#### Build
```bash
make
```

#### Queryring

When deployed with Promethus server, the following PromQL queries might be helpful to get webserver operational status:

###### Average process memory use:
`avg_over_time(go_memstats_alloc_bytes[1m])`

###### Request rate:
`sum(rate(http_requests_total[1m]))`

###### Error rate:
`sum(rate(http_requests_total{status_code!~"2.."}[1m]))`

###### Error ratio (in %):
`sum(rate(http_requests_total{status_code!~"2.."}[1m])) / sum(rate(http_requests_total[1m])) * 100`

###### Average request duration (grouped by status code):
`sum(rate(request_duration_seconds_sum[1m])) by (status_code) / sum(rate(request_duration_seconds_count[1m])) by (status_code)`

###### 85-percentile of request duration (in sec.):
`histogram_quantile(0.85, sum(rate(request_duration_seconds_bucket[1m]))by(le))`

###### 99-percentile of successful request duration (in sec.):
`histogram_quantile(0.99, sum(rate(request_duration_seconds_bucket{status_code="200"}[1m]))by(le))`

###### Average response payload size (in bytes):
`sum(rate(response_size_bytes_sum[1m])) / sum(rate(response_size_bytes[1m]))`

###### 90-percentile of average response payload size (in bytes) for successful requests:
`avg_over_time(response_size_bytes{quantile="0.9", status_code="200"}[1m])`
 