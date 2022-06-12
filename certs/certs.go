// package certs provides some embeded dumb certs.
package certs

import (
	"crypto/tls"
	_ "embed"
)

//go:embed ca.pem
var CA_CERT []byte

//go:embed cakey.pem
var CA_KEY []byte

// Gets a new X509Cert from the embeded cert.
// or panics
func NewX509() tls.Certificate {
	cert, err := tls.X509KeyPair(CA_CERT, CA_KEY)
	if err != nil {
		panic(err)
	}
	return cert
}
