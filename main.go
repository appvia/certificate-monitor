package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/appvia/certificate-monitor/monitor"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

var (
	kubeconfigPath     = flag.String("kubeconfig", "", "Path to kubeconfig file if running outside the Kubernetes cluster")
	ingressDomains     = flag.Bool("ingressNamespaces", false, "Should ingress resources be queried for domains")
	includedDomains    = flag.String("domains", "", "Comma-separated SNI domains to query, or use the MONITOR_DOMAINS environment variable")
	excludedDomains    = flag.String("ignoredDomains", "", "Comma-separated list of domains to exclude from the discovered set, or use the MONITOR_EXCLUDED_DOMAINS environment variable")
	metricsPort        = flag.Int("metricsPort", 8080, "TCP port that the Prometheus metrics listener should use")
	insecureSkipVerify = flag.Bool("insecure", true, "If true, then the InsecureSkipVerify option will be used with the TLS connection, and the remote certificate and hostname will be trusted without verification")
	connectionTimeout  = flag.Int("connectionTimeout", 4, "Timeout when fetching a domain's tls certificate")
)

func loadDomainList(cli_val, env_var string) []string {
	if cli_val == "" {
		return strings.Split(os.Getenv(env_var), ",")
	} else {
		return strings.Split(cli_val, ",")
	}
}

func newLogger() *logrus.Logger {
	var logger = logrus.New()
	logger.Out = os.Stderr
	jsonFormatter := new(logrus.JSONFormatter)
	jsonFormatter.TimestampFormat = time.RFC3339Nano
	logger.Formatter = jsonFormatter
	logger.Level = logrus.InfoLevel
	return logger
}

func main() {
	flag.Parse()

	logger := newLogger()

	var cli *kubernetes.Clientset
	if *ingressDomains {
		kubeClient, err := monitor.NewClientSet(*kubeconfigPath)
		if err != nil {
			logger.Fatalf("Error creating Kubernetes client, exiting: %v", err)
		}
		cli = kubeClient
	}

	mon := &monitor.CertExpiryMonitor{
		Logger:            logger,
		IncludedDomains:   loadDomainList(*includedDomains, "MONITOR_DOMAINS"),
		ExcludedDomains:   loadDomainList(*excludedDomains, "MONITOR_EXCLUDED_DOMAINS"),
		IngressDomains:    *ingressDomains,
		Port:              *metricsPort,
		KubeClient:        cli,
		ConnectionTimeout: *connectionTimeout,
	}

	mon.RunForever()
}
