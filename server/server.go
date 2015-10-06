package server

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/gob"
	"github.com/gnicod/aion/scheduler"
	"log"
	"net"
)

type Server struct {
	l   net.Listener
	sch scheduler.Scheduler
}

func NewServer(sch scheduler.Scheduler) Server {
	// TODO path need to be read from a config
	cert, err := tls.LoadX509KeyPair("/home/ovski/certs/server.pem", "/home/ovski/certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader
	service := "0.0.0.0:8000"
	l, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	serv := Server{l, sch}
	return serv
}

func (s *Server) Listen() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		tlscon, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
			state := tlscon.ConnectionState()
			for _, v := range state.PeerCertificates {
				log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
			}
		}
		go s.serve(conn)
	}
}

func (s *Server) serve(c net.Conn) {
	dec := gob.NewDecoder(c)
	t := &scheduler.Task{}
	dec.Decode(t)
	println("Received : %+v", t)

	println("Command:", string(t.Command))
	println("Expression:", string(t.Expression))
	_, err := c.Write([]byte(t.Command))
	if err != nil {
		log.Fatal("Write: ", err)
	}
}
