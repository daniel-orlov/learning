package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	ConnectToDB()
	security()
}

func security() {
	certificates, err := tls.LoadX509KeyPair(os.Getenv("CERT"), os.Getenv("KEY"))
	if err != nil {
		err = fmt.Errorf("unable to load certificates: %w", err)
		_, _ = fmt.Fprint(os.Stderr, err)
	}
	tlsConf := tls.Config{Certificates: []tls.Certificate{certificates}}
	tls.Listen("tcp", PORT, &tlsConf)
}
