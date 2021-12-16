# Certificate Monitor

Monitors the expiry time of tls certificates and exports prometheus metrics. Target domains can be automatically discovered via ingress resources or specified on the command line/via environment variables

## Usage
```
Usage of ./certificate-monitor:
  -connectionTimeout int
    	Timeout when fetching a domain's tls certificate (default 4)
  -domains string
    	Comma-separated SNI domains to query, or use the MONITOR_DOMAINS environment variable
  -ignoredDomains string
    	Comma-separated list of domains to exclude from the discovered set, or use the MONITOR_EXCLUDED_DOMAINS environment variable
  -ingressNamespaces
    	Should ingress resources be queried for domains
  -insecure
    	If true, then the InsecureSkipVerify option will be used with the TLS connection, and the remote certificate and hostname will be trusted without verification (default true)
  -kubeconfig string
    	Path to kubeconfig file if running outside the Kubernetes cluster
  -metricsPort int
    	TCP port that the Prometheus metrics listener should use (default 8080)
```