package observability

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	duration *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "payment-service",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "method"}),
	}

	reg.MustRegister(m.duration)
	return m
}
