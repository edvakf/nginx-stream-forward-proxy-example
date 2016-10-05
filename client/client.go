package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func req(server string) {
	var localDial = func(network, addr string) (net.Conn, error) {
		return (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial(network, server)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: localDial, //replace Dial
			//DialTLS: localDialTLS, //replace DialTLS
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Get("https://server:443/")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}

func main() {
	//var servers = []string{"server:18080", "proxy1:18080", "proxy2:18080"}
	var servers = []string{"proxy1:18080", "proxy2:18080"}
	//var servers = []string{"server:443"}
	for i := 0; i < 60; i++ {
		req(servers[i%len(servers)])
		time.Sleep(1 * time.Second)
	}
}
