package main

import "github.com/prometheus/client_golang/prometheus"

var PORT = ":1234"
var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "Ali",
		Name:      "my_counter",
		Help:      "This is my counter ",
	})

var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "Ali",
		Name:      "my_gauge",
		Help:      "This is my gauge ",
	})
var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "Ali",
		Name:      "my_histogram",
		Help:      "This is my histogram ",
	})
var summery = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Namespace: "Ali",
		Name:      "my_summary",
		Help:      "This is my summary ",
	})

func main() {

}
