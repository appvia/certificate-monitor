package monitor

import (
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricServer(port int) (*healthHandler, error) {
	hh := &healthHandler{
		healthy: false,
	}
	err := runHTTPListener(hh, port)
	if err != nil {
		return nil, err
	}
	return hh, nil
}

func runHTTPListener(hh *healthHandler, port int) error {
	m := http.NewServeMux()
	m.Handle("/healthz", hh)
	m.Handle("/metrics", promhttp.Handler())
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	go http.Serve(lis, m)
	return nil
}

type healthHandler struct {
	healthy bool
}

func (hh *healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hh.healthy {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Unhealthy"))
}
