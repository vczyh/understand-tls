package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
)

const serverPort = 9999

func main() {
	config, err := tlsConfig()
	if err != nil {
		log.Printf("tlsConfig(): %v", err)
		return
	}

	l, err := tls.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", serverPort), config)
	if err != nil {
		log.Printf("net.Listen(): %v\n", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("l.Accept(): %v\n", err)
			continue
		}
		log.Println("New Connection")

		go handleConnection(conn)
	}
}

func tlsConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
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
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return config, nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Printf("r.ReadString(): %v\n", err)
			return
		}
		log.Printf("Receive message: %s\n", msg)

		_, err = conn.Write([]byte("ACK\n"))
		if err != nil {
			log.Printf("conn.Write(): %v\n", err)
			return
		}
	}
}
