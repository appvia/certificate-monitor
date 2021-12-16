package monitor

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CertExpiryMonitor struct {
	Logger            *logrus.Logger
	IncludedDomains   []string
	ExcludedDomains   []string
	IngressDomains    bool
	Port              int
	KubeClient        *kubernetes.Clientset
	ConnectionTimeout int
}

func (m *CertExpiryMonitor) GetIngressDomains() ([]string, error) {
	out := []string{}

	if !m.IngressDomains {
		return out, nil
	}
	listOptions := metav1.ListOptions{}
	extApi := m.KubeClient.NetworkingV1()

	ings, err := extApi.Ingresses("").List(context.Background(), listOptions)
	if err != nil {
		return nil, err
	}

	for _, ingress := range ings.Items {
		for tls_block_index := range ingress.Spec.TLS {
			for host_index := range ingress.Spec.TLS[tls_block_index].Hosts {
				out = append(out, ingress.Spec.TLS[tls_block_index].Hosts[host_index])
			}
		}
	}
	return out, nil
}

func (m *CertExpiryMonitor) GetTargetDomains() ([]string, error) {
	IngressDomains, err := m.GetIngressDomains()
	if err != nil {
		return nil, err
	}
	return stringSetSubtract(stringSetUnion(IngressDomains, m.IncludedDomains), m.ExcludedDomains), nil
}

func (m *CertExpiryMonitor) RunForever() {
	m.Logger.Infof("Starting the webserver on port %v", m.Port)
	hh, err := StartMetricServer(m.Port)
	if err != nil {
		m.Logger.Fatalf("Error when starting the webserver: %v", err)
	}
	hh.healthy = true
	m.Logger.Infof("Server is started, serving metrics")
	for {
		domains, err := m.GetTargetDomains()
		if err != nil {
			m.Logger.Errorf("Error when fetching target domains: %v", err)
			continue
		}
		for d := range domains {
			cert, err := GetCertificate(domains[d], m.ConnectionTimeout)
			if err != nil {
				m.Logger.Errorf("Error when fetching certificate: %v", err)
				continue
			}
			certificateSecondsSinceIssued.WithLabelValues(domains[d]).Set(SecondsSinceIssued(cert))
			certificateSecondsUntilExpiry.WithLabelValues(domains[d]).Set(SecondsUntilExpiry(cert))
		}
		time.Sleep(10 * time.Second)
	}
}
