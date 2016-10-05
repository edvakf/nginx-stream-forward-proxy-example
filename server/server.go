package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	certFile, _ := filepath.Abs("oreore.crt")
	keyFile, _ := filepath.Abs("oreore.key")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("connected from %s (protocol %s)\n", r.RemoteAddr, r.Proto)
		fmt.Fprintf(w, "connected from %s (protocol %s)\n", r.RemoteAddr, r.Proto)
	})

	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
