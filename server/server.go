package server

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
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
		_, ok := conn.(*tls.Conn)
		if ok {
			log.Print("ok=true")
		}
		go s.serve(conn)
	}
}

func (s *Server) serve(c net.Conn) {
	dec := gob.NewDecoder(c)
	//t := &scheduler.Task{}
	t := Command{}
	dec.Decode(t)

	//s.sch.AddTask(t)

	//println("Command:", string(t.Command))
	//println("Expression:", string(t.Expression))
	res := Response{Content: "list"}
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(res)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	_, err = c.Write(network.Bytes())
	defer c.Close()
	if err != nil {
		log.Fatal("write error:", err)
	}
}
