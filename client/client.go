package client

import (
	"bytes"
	"crypto/tls"
	//"crypto/x509"
	//"fmt"
	"log"
	"net"
	//"os"
	"encoding/gob"
	"github.com/gnicod/aion/scheduler"
	"time"
)

type Client struct {
	conn net.Conn
}

func NewClient() Client {
	cert, err := tls.LoadX509KeyPair("/home/ovski/certs/client.pem", "/home/ovski/certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "127.0.0.1:8000", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	log.Println("client: connected to: ", conn.RemoteAddr())

	/*
		state := conn.ConnectionState()
			for _, v := range state.PeerCertificates {
				fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
				fmt.Println(v.Subject)
			}
	*/
	client := Client{conn}
	go client.reader()
	return client
}

func (c *Client) reader() {
	reply := make([]byte, 256)
	n, err := c.conn.Read(reply)
	if err != nil {
		log.Fatal("write error:", err)
	}
	log.Printf("client: got %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
}

func (c *Client) AddTask(t scheduler.Task) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(t)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	_, err = c.conn.Write(network.Bytes())
	defer c.conn.Close()
	if err != nil {
		log.Fatal("add task write error:", err)
	}
}

func (c *Client) Send() {
	defer c.conn.Close()
	for {
		_, err := c.conn.Write([]byte("hi"))
		if err != nil {
			log.Fatal("write error:", err)
			break
		}
		time.Sleep(10e9)
	}
}
