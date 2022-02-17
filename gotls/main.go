package main

import "log"

func main() {
	if err := GenerateCA(); err != nil {
		log.Fatalf("GenerateCA(): %v", err)
	}

	if err := GenerateServer([]string{"10.0.44.59"}); err != nil {
		log.Fatalf("GenerateServer(): %v", err)
	}

	if err := GenerateClient([]string{"10.0.8.13"}); err != nil {
		log.Fatalf("GenerateClient(): %v", err)
	}
}
