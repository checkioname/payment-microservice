package observability

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	duration *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "payment_service",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "method"}),
	}

	reg.MustRegister(m.duration)
	return m
}

func (m *Metrics) ObserveDuration(status, method string, duration float64) {
	m.duration.WithLabelValues(status, method).Observe(duration)
}
