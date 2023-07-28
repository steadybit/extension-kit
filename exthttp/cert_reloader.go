package exthttp

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"
)

type CertReloader struct {
	CertFile          string
	KeyFile           string
	cachedCert        *tls.Certificate
	cachedCertModTime time.Time
}

func NewCertReloader(certFile, keyFile string) *CertReloader {
	return &CertReloader{CertFile: certFile, KeyFile: keyFile}
}

func (cr *CertReloader) GetCertificate(_ *tls.ClientHelloInfo) (*tls.Certificate, error) {
	stat, err := os.Stat(cr.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed checking key file modification time: %w", err)
	}

	if cr.cachedCert == nil || stat.ModTime().After(cr.cachedCertModTime) {
		pair, err := tls.LoadX509KeyPair(cr.CertFile, cr.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed loading tls key pair: %w", err)
		}

		cr.cachedCert = &pair
		cr.cachedCertModTime = stat.ModTime()
	}

	return cr.cachedCert, nil
}
