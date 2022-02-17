package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"time"
)

func GenerateServer(hosts []string) error {
	// CA certificate
	caCertBytes, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		return err
	}
	caCertBlock, _ := pem.Decode(caCertBytes)
	if caCertBlock == nil {
		return fmt.Errorf("parse ca cert failed")
	}
	caCert, err := x509.ParseCertificate(caCertBlock.Bytes)
	if err != nil {
		return err
	}

	// CA private key
	caPrivateKeyBytes, err := ioutil.ReadFile("certs/ca.key")
	if err != nil {
		return err
	}
	caPrivateKeyBlock, _ := pem.Decode(caPrivateKeyBytes)
	if caPrivateKeyBlock == nil {
		return fmt.Errorf("parse ca private key failed")
	}
	caPrivateKey, err := x509.ParsePKCS8PrivateKey(caPrivateKeyBlock.Bytes)
	if err != nil {
		return err
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	keyUsage :=
		x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageDataEncipherment

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(time.Hour * 24 * 365 * 10)

	template := x509.Certificate{
		PublicKey:    privateKey.PublicKey,
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"ZhangYuheng Inc"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:    keyUsage,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},

		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, caCert, publicKey(privateKey), caPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to create certificate: %v", err)
	}

	certOut, err := os.Create("certs/server.crt")
	if err != nil {
		return fmt.Errorf("failed to open cert.pem for writing: %v", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return fmt.Errorf("failed to write data to cert.pem: %v", err)
	}
	if err := certOut.Close(); err != nil {
		return fmt.Errorf("error closing cert.pem: %v", err)
	}

	keyOut, err := os.OpenFile("certs/server.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open key.pem for writing: %v", err)
	}
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("unable to marshal private key: %v", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privateKeyBytes}); err != nil {
		return fmt.Errorf("failed to write data to key.pem: %v", err)
	}
	if err := keyOut.Close(); err != nil {
		return fmt.Errorf("error closing key.pem: %v", err)
	}

	return nil
}
