package monitor

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"time"
)

func GetCertificate(domain string, timeout int) (*x509.Certificate, error) {
	dialer := new(net.Dialer)
	dialer.Timeout = time.Duration(timeout) * time.Second
	conn, err := tls.DialWithDialer(dialer, "tcp", fmt.Sprintf("%s:443", domain), &tls.Config{})

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn.ConnectionState().PeerCertificates[0], nil
}

func SecondsUntilExpiry(cert *x509.Certificate) float64 {
	return time.Until(cert.NotAfter).Seconds()
}

func SecondsSinceIssued(cert *x509.Certificate) float64 {
	return time.Since(cert.NotBefore).Seconds()
}
