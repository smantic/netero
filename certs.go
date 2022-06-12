package goproxy

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
)

func init() {
	if goproxyCaErr != nil {
		panic("Error parsing builtin CA " + goproxyCaErr.Error())
	}
	var err error
	if GoproxyCa.Leaf, err = x509.ParseCertificate(GoproxyCa.Certificate[0]); err != nil {
		panic("Error parsing builtin CA " + err.Error())
	}
}

var tlsClientSkipVerify = &tls.Config{InsecureSkipVerify: true}

var defaultTLSConfig = &tls.Config{
	InsecureSkipVerify: true,
}

//go:embed ca.pem
var CA_CERT []byte

//go:embed cakey.pem
var CA_KEY []byte

var GoproxyCa, goproxyCaErr = tls.X509KeyPair(CA_CERT, CA_KEY)
