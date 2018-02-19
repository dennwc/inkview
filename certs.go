package ink

import (
	"crypto/x509"
	"encoding/asn1"
)

// InitCerts will read system certificates pool.
//
// This pool is usually populated by the first call to tls.Dial or similar,
// but this operation might take up to 30 sec on some devices, leading to handshake timeout.
//
// Calling this function before dialing will fix the problem.
func InitCerts() error {
	// hand-crafted fake cert that will force system pool to be populated
	// but will fail with an error directly after this
	cert := x509.Certificate{
		Raw: []byte{0},
		UnhandledCriticalExtensions: []asn1.ObjectIdentifier{nil},
	}
	_, err := cert.Verify(x509.VerifyOptions{})
	if _, ok := err.(x509.SystemRootsError); ok {
		return err
	}
	return nil
}
