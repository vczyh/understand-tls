package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

const (
	serverAddr = "10.0.44.59"
	serverPort = 9999
)

func main() {
	config, err := tlsConfig()
	if err != nil {
		log.Printf("tlsConfig(): %v", err)
		return
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, serverPort), config)
	if err != nil {
		log.Printf("net.Dial(): %v\n", err)
		return
	}

	_, err = conn.Write([]byte("I am Client\n"))
	if err != nil {
		log.Printf("conn.Write(): %v\n", err)
		return
	}

	r := bufio.NewReader(conn)
	msg, err := r.ReadString('\n')
	if err != nil {
		log.Printf("r.ReadString(): %v\n", err)
		return
	}
	log.Printf("Receive message: %s\n", msg)
}

func tlsConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("certs/client.crt", "certs/client.key")
	if err != nil {
		return nil, fmt.Errorf("tls.LoadX509KeyPair(): %v", err)
	}

	caCertBytes, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(): %v", err)
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCertBytes); !ok {
		return nil, fmt.Errorf("certPool.AppendCertsFromPEM(): %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	return config, nil
}
