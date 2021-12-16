package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	certificateSecondsSinceIssued = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certificate_expiry_monitor",
			Name:      "seconds_since_cert_issued",
			Help:      "Secods since the certificate was issued",
		},
		[]string{"domain"},
	)
	certificateSecondsUntilExpiry = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certificate_expiry_monitor",
			Name:      "seconds_until_cert_expiry",
			Help:      "Seconds until the certificate expires",
		},
		[]string{"domain"},
	)
)
