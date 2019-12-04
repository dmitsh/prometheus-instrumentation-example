package main

import (
	"log"
	"net/http"

	"github.com/dmitsh/prometheus-instrumentation-example/internal/sim"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	gauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_percent_used",
			Help: "CPU percent used.",
		},
		[]string{"host", "module", "set"})

	counter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"path", "status_code"})

	histogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration_seconds",
			Help:    "Time (in seconds) spent serving HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "status_code"})

	summary = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "response_size_bytes",
			Help:       "Response size in bytes.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path", "status_code"})
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// simulate response
	status := sim.GenerateStatusCode()
	reqTime := sim.GenerateRequestTime()
	respSize := sim.GenerateResponseSize()

	// update metrics
	//gauge.WithLabelValues(host, module, set).Set(cpu)
	counter.WithLabelValues(r.URL.Path, status).Inc()
	histogram.WithLabelValues(r.URL.Path, status).Observe(reqTime)
	summary.WithLabelValues(r.URL.Path, status).Observe(respSize)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
