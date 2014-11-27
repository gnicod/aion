package server

import (
    "log"
    "net"
    "encoding/gob"
    "github.com/gnicod/aion/scheduler"
)

type Server struct {
    l net.Listener
    sch scheduler.Scheduler
}

func NewServer(sch scheduler.Scheduler) Server{
    l, err := net.Listen("unix","/tmp/aion.sock" )
    if err != nil {
        log.Fatal("listen error:", err)
    }
    serv := Server{l, sch}
    return serv
}


func (s *Server) Listen(){
    for {
        fd, err := s.l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }
        go s.serve(fd)
    }
}

func (s *Server) serve(c net.Conn){
        dec := gob.NewDecoder(c)
        t := &scheduler.Task{}
        dec.Decode(t)
        println("Received : %+v", t);

        println("Command:", string(t.Command))
        println("Expression:", string(t.Expression))
        _, err := c.Write([]byte(t.Command))
        if err != nil {
            log.Fatal("Write: ", err)
        }
}
